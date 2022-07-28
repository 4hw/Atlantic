package meta

type Mirai struct {
        AttackPrefix string `json:"AttackPrefix"`
        CurrentCNC string `json:"CurrentCNC"`
        Username string `json:"Username"`
        Password string `json:"Password"`

        Methods []struct {
                Name string `json:"Name"`
                Description string `json:"Description"`
                Args []string `json:"Args"`
        } `json:"Methods"`
}


type Slaves struct {
        AttackPrefix string `json:"AttackPrefix"`
        Slaves []struct {
                ID        uint16   `json:"ID"`
                Name  string   `json:"Name"`
                Description string `json:"Description"`
                DefaultPort string `json:"DefaultPort"`
                Admin bool `json:"Admin"`
                Vip bool `json:"Vip"`
        } `json:"Slaves"`
}

