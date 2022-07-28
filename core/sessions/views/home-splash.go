package views

import (
	"strings"

	"golang.org/x/crypto/ssh"

	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/sessions/sessions"
	"atlantic-cnc/core/models/client/terminal"
)

func Home_Splash(channel ssh.Channel, conn *ssh.ServerConn, user *database.User, session *sessions.Session_Store) {

	if session.CurrentTheme == nil {
		error,_ := terminal.Banner("home-splash", user, channel, true, false, nil)
		if error != nil {
			return
		}
	} else {
		error,_ := terminal.Banner(strings.Split(session.CurrentTheme.Views_HomeClear, "/")[1], user, channel, true, false, nil)
		if error != nil {
			return
		}
	}



	Prompt(session)

}