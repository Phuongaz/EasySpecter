package commands

import (
	"log"
	"strconv"
	"strings"

	"github.com/phuongaz/easyspecter/specter"
)

type Specter struct{}

func (Specter) Execute(parameters []string, l *log.Logger) {
	if len(parameters) == 0 {
		l.Println("Command: specter <type> <parameters>")
		l.Println("Types & parameters:")
		l.Println("  + move <name> <x>, <y>, <z>")
		l.Println("  + chat <name> <message>")
		l.Println("  + quit <name>")
		ReadConsole(l)
		return
	}

	if parameters[0] == "move" {
		spt, ok := specter.GetSpecter(parameters[1])
		if !ok {
			l.Fatal("Specter not found")
			return
		}
		x, _ := strconv.ParseFloat(parameters[2], 32)
		y, _ := strconv.ParseFloat(parameters[3], 32)
		z, _ := strconv.ParseFloat(parameters[4], 32)
		spt.Move(float32(x), float32(y), float32(z))
		l.Printf("Specter %s moved to %f, %f, %f", spt.Conn.IdentityData().DisplayName, x, y, z)
	}

	if parameters[0] == "chat" {
		spt, ok := specter.GetSpecter(parameters[1])
		if !ok {
			l.Fatal("Specter not found")
			return
		}
		mess := strings.Join(parameters[2:], " ")
		spt.Chat(mess)
		l.Printf("Specter %s chat: %s", spt.Conn.IdentityData().DisplayName, mess)
	}

	if parameters[0] == "quit" {
		spt, ok := specter.GetSpecter(parameters[1])
		if !ok {
			l.Fatal("Specter not found")
			return
		}
		spt.Quit()
		l.Printf("Specter %s quit", spt.Conn.IdentityData().DisplayName)
	}

	if parameters[0] == "list" {
		l.Println("Specters:")
		for _, spt := range specter.GetSpecters() {
			l.Println(spt.Conn.IdentityData().DisplayName + " in " + spt.Address)
		}
	}
}
