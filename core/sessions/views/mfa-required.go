package views

import (
        "atlantic-cnc/core/mysql"

        "errors"
        "github.com/xlzd/gotp"
        "golang.org/x/crypto/ssh"
        "golang.org/x/term"
)

func MFARequire(channel ssh.Channel, conn *ssh.ServerConn, User *database.User) error {
        channel.Write([]byte("                __  ____________       __________  ____  ______\r\n"))
        channel.Write([]byte("               /  |/  / ____/   |     / ____/ __ \\/ __ \\/ ____/\r\n"))
        channel.Write([]byte("              / /|_/ / /_  / /| |    / /   / / / / / / / __/\r\n"))
        channel.Write([]byte("             / /  / / __/ / ___ |   / /___/ /_/ / /_/ / /___\r\n"))
        channel.Write([]byte("            /_/  /_/_/   /_/  |_|   \\____/\\____/_____/_____/\r\n\n\n"))
        channel.Write([]byte("\x1b[0m                       MFA Code> "))
        Reader := term.NewTerminal(channel, "\033[7;34H")
        Code, err := Reader.ReadLine()
        if err != nil {
                channel.Close()
                return err
        }

        TOTP := gotp.NewDefaultTOTP(User.MFA)
        if TOTP.Now() != Code {
                channel.Write([]byte("                 LOG: Invaild MFA Code!\r\n"))
                channel.Close()
                return errors.New("Invaild MFA Code")
        }

        return nil
}
