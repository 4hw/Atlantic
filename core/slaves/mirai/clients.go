package slaves

import (
        "crypto/rand"
        "encoding/binary"
        "log"
        "net"
        "strings"
        "sync"
)

var (
        clients *Clients = &Clients{
                all: make(map[uint64]*Client),
        }

        queueLen = 30
)

type Clients struct {
        mutex sync.RWMutex

        all map[uint64]*Client

        count int
}

type Client struct {
        ID uint64

        Conn net.Conn
        Name string

        Queue chan []byte

        parent *Clients
}

func (c *Clients) Add(client *Client) error {
        c.mutex.Lock()
        defer c.mutex.Unlock()

        client.Queue = make(chan []byte, queueLen)

        buf := make([]byte, 8)
        for {
                if _, err := rand.Read(buf); err != nil {
                        return err
                }

                id := binary.BigEndian.Uint64(buf)
                if _, ok := c.all[id]; ok == false {
                        client.ID = id
                        c.all[id] = client
                        c.count++
                        return nil
                }
        }

}

func (c *Client) Remove() {
        c.parent.mutex.Lock()
        defer c.parent.mutex.Unlock()

        close(c.Queue)
        c.parent.count--

        delete(c.parent.all, c.ID)
}

func Count() int {
        clients.mutex.RLock()
        defer clients.mutex.RUnlock()

        return clients.count
}

func Send(payload []byte) {
        clients.mutex.RLock()
        defer clients.mutex.RUnlock()

        for _, client := range clients.all {
                if len(client.Queue) >= queueLen-1 {
                        continue
                }

                client.Queue <- payload
        }
}

func Clone() []Client {
        clients.mutex.RLock()
        defer clients.mutex.RUnlock()

        var list []Client

        for _, client := range clients.all {
                list = append(list, *client)
        }

        return list
}

func (c *Client) IP() string {
        c.parent.mutex.Lock()
        defer c.parent.mutex.Unlock()

        ip, _, err := net.SplitHostPort(c.Conn.RemoteAddr().String())
        if err != nil {
                log.Println(" [SLAVES] ["+err.Error()+"]")
                return c.Conn.RemoteAddr().String()
        }

        return ip
}

func (c *Client) CleanName() string {

        name := strings.Replace(c.Name, "\n", "\\n", -1)
        name = strings.Replace(name, "\r", "\\r", -1)
        name = strings.Replace(name, "\t", " ", -1)
        name = strings.Replace(name, "\x1b", `\x1b`, -1)

        return name
}
