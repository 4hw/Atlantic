package build

import (
        "atlantic-cnc/core/models/json/meta"
        "atlantic-cnc/core/models/versions"

        "encoding/json"
        "io/ioutil"
        "log"
        "os"
)


var AttackSyncs *meta.Attack
var AttackAPI *meta.AttackMethod

func LoadAttacks() (bool, error) {
        File, errors := os.Open("assets/attack.json"); if errors != nil {
                log.Println("Failed To Parse \"attack.json\"",errors)
                err := File.Close(); if err != nil {
                    log.Println("Failed Closing Of \"attack.json\"")
                    return false,err
                }
                return false,errors
        }
        byteValue, errors := ioutil.ReadAll(File); if errors != nil {
                log.Println("Failed To Read \"attack.json\"",errors)

                err := File.Close(); if err != nil {
                    log.Println("Failed Closing Of \"attack.json\"")
                    return false,err
                }
                return false,errors
        }
        var CG meta.Attack
        errors = json.Unmarshal(byteValue, &CG)
        if errors != nil {
                log.Println("Failed To Parse \"attack.json\"", errors)

                err := File.Close(); if err != nil {
                    log.Println("Failed Closing Of \"attack.json\"")
                    return false,err
                }
                return false,errors
        }

        err := File.Close(); if err != nil {
                log.Println("Failed Closing Of \"attack.json\"")
                return false,err
        }


        AttackSyncs = &CG
        return true, nil
}

func NewParse_Attack_Json() (error) {
        File, error := os.Open(versions.GOOS_Edition.Make["API_Attack"])
        if error != nil {
                return error
        }

        defer File.Close()

        ByteVal, error := ioutil.ReadAll(File)
        if error != nil {
                return error
        }

        var NewConfigAllo meta.AttackMethod
        error = json.Unmarshal(ByteVal, &NewConfigAllo)
        if error != nil {
                return error
        }

        AttackAPI = &NewConfigAllo


        return nil
}
