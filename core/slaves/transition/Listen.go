package transition

import (
        "atlantic-cnc/core/models/json/build"
        "log"
        "net"
)

func Connection() {
        Nets, error := net.Dial("tcp", build.Option.SlaveTransition.LoopbackPort); if error != nil {
                log.Printf(" [Failed to use port] [%s]", build.Option.SlaveTransition.LoopbackPort)
                Connection()
                return
        } else {
                log.Printf(" [Connected to Mirai server] [%s]", build.Option.SlaveTransition.LoopbackPort)
                Sort(Nets)
                return
        }
}
