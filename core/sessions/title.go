package masters

import (
	"time"
	"atlantic-cnc/core/sessions/sessions"
	"atlantic-cnc/core/models/client/terminal"
	"atlantic-cnc/core/models/json/build"
)


func TitleWorker() {

	for {


		if build.Config.Extra.TitleSpinner {
			for I := 0; I < len(build.Config.Extra.TitleSpinnerFrames); I++ {
				for _, session := range sessions.SessionMap {

					var Map = map[string]string {
						"spinner" : build.Config.Extra.TitleSpinnerFrames[I],
					}

					error, title := terminal.Banner("title", session.User, session.Channel, false, true, Map)
					if error != nil {
						continue
					}
		
					session.Channel.Write([]byte("\033]0;"+title+"\007"))
					continue
				}
		
				time.Sleep(1 * time.Second)
			}
		} else {
			for _, session := range sessions.SessionMap {
				error, title := terminal.Banner("title", session.User, session.Channel, false, true, nil)
				if error != nil {
					continue
				}
	
				session.Channel.Write([]byte("\033]0;"+title+"\007"))
				continue
			}
	
			time.Sleep(1 * time.Second)
		}
	}
}