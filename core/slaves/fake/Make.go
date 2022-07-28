package Fake

import (
        "atlantic-cnc/core/models/json/build"
        "math/rand"
)

func MakeLoop() int {
        Num := rand.Intn(build.Option.FakeSlaves.MaxGen-build.Option.FakeSlaves.MinGen) + build.Option.FakeSlaves.MinGen
        return Num
}
