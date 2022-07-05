package commands

import (
	"log"

	"github.com/phuongaz/easyspecter/specter"
)

type Join struct{}

func (Join) Execute(parameters []string, l *log.Logger) {
	if len(parameters) == 0 {
		l.Fatal("Command: join <name> <ip:port>")
		ReadConsole(l)
		return
	}
	spt := &specter.SpecterNormal{
		Specter: specter.Specter{
			Log:     l,
			Address: parameters[1],
		},
		Name: parameters[0],
	}
	_, err := spt.Login()
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	specter.AddSpecter(&spt.Specter)
}
