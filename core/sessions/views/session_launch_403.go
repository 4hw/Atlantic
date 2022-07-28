package views

import (
	"time"
	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/models/client/terminal"

	"golang.org/x/crypto/ssh"
)


func Sessions_launch_403(channel ssh.Channel, conn *ssh.ServerConn, User *database.User) error {

	error,_ := terminal.Banner("sessions_launch-403", User, channel, true, false, nil)
	if error != nil {
		return error
	}

	time.Sleep(10 * time.Second)

	return nil
}