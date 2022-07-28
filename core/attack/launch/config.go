package attack_launch

type Method struct {
        Method                  string
        APIMethod                       string
        DefaultPort             string

        AdminMethod             bool
        VIPMethod                       bool

        API                             string

        LimitMaxTime bool
        LimitMax int
}
