package util_Command

import (
        "atlantic-cnc/core/sessions/sessions"
)

func init() {

	Register(&Command{
		Name: "exit",

		Description: "exit out of cnc",

		Admin: false,
		Reseller: false,
		Vip: false,

                Execute: func(Session *sessions.Session_Store, cmd []string) error {
                        Session.Channel.Close()
                        return nil
		},
	})
}
