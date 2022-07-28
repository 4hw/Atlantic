package meta

type Options struct {
        FakeSlaves struct {
                Status bool `json:"Status"`
                MinGen int `json:"MinGenSet"`
                MaxGen int `json:"MaxGenSet"`
        } `json:"FakeSlaves"`

        SlaveTransition struct {
                Status bool `json:"Status"`
                LoopbackPort string `json:"LoopbackPort"`
        } `json:"SlaveTransition"`
}
