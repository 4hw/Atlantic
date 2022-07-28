package views

import (
        "errors"
        "fmt"
        "net/url"
        "strings"

        "atlantic-cnc/core/mysql"
        "atlantic-cnc/core/models/json/build"
        "atlantic-cnc/core/functions"
        "github.com/xlzd/gotp"
        "golang.org/x/crypto/ssh"
        "golang.org/x/term"
)

func CreateMFA(channel ssh.Channel, conn *ssh.ServerConn) error {
        secret := functions.GenTOTPSecret()
        totp := gotp.NewDefaultTOTP(secret)
        qrs := functions.New()

        qrcode := qrs.Get("otpauth://totp/" + url.QueryEscape(build.Config.AppConfig.AppName) + ":" + url.QueryEscape(conn.User()) + "?secret=" + secret + "&issuer=" + url.QueryEscape(build.Config.AppConfig.AppName) + "&digits=6&period=30").Sprint()
        fmt.Fprintln(channel, strings.ReplaceAll(qrcode, "\n", "\r\n"))
        fmt.Fprint(channel, "You may scan this code to register your account info a 2FA App, Google Auth, Twilio Authy\r\n")
        fmt.Fprint(channel, "or enter this Code: "+secret+"\r\n")
        term := term.NewTerminal(channel, "Code> ")

        Code, error := term.ReadLine(); if error != nil {
                fmt.Fprint(channel, "\r\n")
                return errors.New("KILLING")
        }

        if totp.Now() != Code {
                fmt.Fprintln(channel, "Invaild MFA Code")
                return errors.New("KILLING")
        }

        errors := database.EditFeild(conn.User(), "MFA", secret); if !errors {
                fmt.Fprintln(channel, "Failed to enable MFA!", error)
        } else {
                channel.Write([]byte(""))
        }

        return nil
}
