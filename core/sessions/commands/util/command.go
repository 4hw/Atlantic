package util_Command

import (
	"log"
	"sync"
	"atlantic-cnc/core/sessions/sessions"
	"atlantic-cnc/core/models/versions"
)


var (
	Commands = make(map[string]*Command)
	NyxMap   sync.Mutex
)

type Command struct {
	Name string
	
	Description string

	Admin bool
	Reseller bool
	Vip bool

	Execute func(Session *sessions.Session_Store, cmd []string) error
}

func Register(command *Command) {

	Command := Commands[command.Name]
	if Command != nil {
		log.Println("[compile issue] [util command "+command.Name+" already exist]")
		return
	}

	NyxMap.Lock()
	Commands[command.Name] = command
	NyxMap.Unlock()

	return
}

func Get(name string) *Command {
	if versions.GOOS_Edition.Util_Commands {
		cmd := Commands[name]
		return cmd
	} else {
		return nil
	}
}