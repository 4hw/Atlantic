package License

import (
        "fmt"
        "io/ioutil"
        "log"
        "time"

        "github.com/Cryptolens/cryptolens-golang/cryptolens"
)


func LicenseGet() bool {
        Key, err := ioutil.ReadFile("assets/license.key"); if err != nil {
                log.Println("Failed to Find your License Key File!")
                return false
        }

        token := "WyIyMTk1ODU1MiIsImJmb1lUL1Z6UU5vOEU0Zkp4NUE5UisxcXlWM1V0QmtBVC8xYUJVM0kiXQ=="
        publicKey := "<RSAKeyValue><Modulus>2/9O03x3iWRWDLOzf/yLyr4eTrjowJWl6C/dzQRDWZkm3e0wCztwex4mvIpIDPVG/G4y3UaUbDUGj5AP57JIO0vh2iQTLwYT3hVkTV9aeqBtWePYfjfNImxewTDN0Tmn7ZYfceUvf5No7ypmbMjfRT4wrJMY+3e7y6ORhPfSuFQvOf/VS503cKl5VAYgwx7c+/D/fMvY55nsqkvBcHLxTRccCcbUS7aUqQ5U0dTgI42bVx4oeMuvee8cOOXcRuUWC8hdMwKBgbRnINJ2ZdH0/NkcAlwKvqdf+4xPszxKzspDAAhMmAvv5RevEkx74q7SEBvrvboZqQhhB5EIozkz6w==</Modulus><Exponent>AQAB</Exponent></RSAKeyValue>"

        licenseKey, err := cryptolens.KeyActivate(token, cryptolens.KeyActivateArguments{
                ProductId:   11343,
                Key:         string(Key),
        })

        if err != nil || !licenseKey.HasValidSignature(publicKey) {
                fmt.Println("License key activation failed!")
                return false
        }

        if time.Now().After(licenseKey.Expires) {
                fmt.Println("License key has expired")
                return false
        }

        if licenseKey.F3 {
                log.Println(" [Livewire DLC is Active on this license]")
                LiveWire = true
        } else {
                LiveWire = false
        }



        return true
}
