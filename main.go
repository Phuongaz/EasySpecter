package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/phuongaz/easyspecter/specter"
	"github.com/phuongaz/easyspecter/xbl"
)

var specters = map[string]*specter.Specter{}

func main() {
	log := log.New(os.Stdout, "[SPECTER] ", log.LstdFlags)
	log.Println("-------EASY SPECTER - AUTHOR PHUONGAZ--------")
	log.Println("1. XBOX ACCOUNT LOG: Xbox account required")
	log.Println("	+ Login command: xbox login")
	log.Println("	+ Spawn command: xbox join <ip:port>")
	log.Println("2. JOIN GAME: No account required")
	log.Println("	+ Must disable xbox-auth in server")
	log.Println("	+ Spawn command: join <ip:port>")
	log.Println("3. LIST: List all specters")
	log.Println("	+ Command: list")
	log.Println("4. QUIT: Quit specter")
	log.Println("	+ Command: quit <specter name>")
	log.Println("5. SPECTER: Specter control")
	log.Println("	+ Chat Command: specter chat <specter name> <message>")
	log.Println("	+ Move Command: specter move <specter name> <x> <y> <z>")
	log.Println("---------------------------------")
	readConsole(log)
}

func readConsole(log *log.Logger) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		if len(text) == 0 {
			continue
		}
		args := strings.Split(text, " ")
		switch args[0] {
		case "xbox":
			if args[1] == "login" {
				log.Println("Login...")
				if err := xbl.InitializeToken(log); err != nil {
					log.Println("Error: ", err)
				}
				log.Println("Login success")
			}
			if args[1] == "join" {
				spt := specter.SpecterXbox{
					Specter: specter.Specter{
						Log:     log,
						Address: args[2],
					},
				}
				_, err := spt.Login(args[2])
				if err != nil {
					log.Println("Error: ", err)
					return
				}
				specters[spt.Specter.Conn.IdentityData().DisplayName] = &spt.Specter
			}
		case "join":
			spt := specter.SpecterNormal{
				Specter: specter.Specter{
					Log:     log,
					Address: args[1],
				},
				Name: "EasySpecter",
			}
			_, err := spt.Login(args[1])
			if err != nil {
				log.Println("Error: ", err)
				return
			}
			specters[spt.Specter.Conn.IdentityData().DisplayName] = &spt.Specter
		case "list":
			log.Println("Specters:")
			for _, spt := range specters {
				log.Println(spt.Conn.IdentityData().DisplayName + " in " + spt.Address)
			}
		case "quit":
			name := args[1]
			if spt, ok := specters[name]; ok {
				spt.Conn.Close()
				delete(specters, name)
				log.Println("Specter " + name + " quit")
			} else {
				log.Println("Specter " + name + " not found")
			}
		case "specter":
			spt := specters[args[1]]
			if spt == nil {
				log.Println("Specter " + args[1] + " not found")
			}
			if args[2] == "chat" {
				message := strings.Join(args[3:], " ")
				spt.Chat(message)
				log.Println(spt.Conn.IdentityData().DisplayName + ": " + message)
			}
			if args[2] == "move" {
				x, _ := strconv.ParseFloat(args[3], 32)
				y, _ := strconv.ParseFloat(args[4], 32)
				z, _ := strconv.ParseFloat(args[5], 32)
				spt.Move(float32(x), float32(y), float32(z))
				log.Println(spt.Conn.IdentityData().DisplayName + " move to " + args[3] + " " + args[4] + " " + args[5])
			}
		}
	}
}
