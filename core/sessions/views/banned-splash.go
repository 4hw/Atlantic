package views

import (
	"time"

	"golang.org/x/crypto/ssh"

	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/models/client/terminal"
)


func LoginBanned(channel ssh.Channel, conn *ssh.ServerConn, user *database.User) {
	error,_ := terminal.Banner("login-banned", user, channel, true, false, nil)
	if error != nil {
		return
	}

	time.Sleep(10 * time.Second)
	channel.Close()
	return
}