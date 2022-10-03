package commands

import (
	"log"

	"github.com/phuongaz/easyspecter/specter"
	"github.com/phuongaz/easyspecter/xbl"
)

type Xbox struct{}

func (Xbox) Execute(parameters []string, l *log.Logger) {
	if len(parameters) == 0 {
		l.Fatal("Command: xbox <login/join [ip:port]>")
		ReadConsole(l)
		return
	}
	if parameters[0] == "login" {
		l.Println("Login...")
		if err := xbl.InitializeToken(l); err != nil {
			log.Println("Error: ", err)
		}
	}
	if parameters[0] == "join" {
		if len(parameters) == 1 {
			l.Fatal("Command: xbox join <ip:port>")
			ReadConsole(l)
			return
		}
		spt := specter.SpecterXbox{
			Specter: specter.Specter{
				Log:     l,
				Address: parameters[1],
			},
		}
		_, err := spt.Login()
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		specter.AddSpecter(&spt.Specter)
		l.Println(spt.Conn.IdentityData().DisplayName + " joined")
	}
}
