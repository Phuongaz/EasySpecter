package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/phuongaz/easyspecter/specter"
	"github.com/phuongaz/easyspecter/xbl"
)

var specters = map[string]*specter.Specter{}

func main() {

	log := log.Default()

	log.Println("-------EASY SPECTER - AUTHOR PHUONGAZ--------")
	log.Println("1. XBOX ACCOUNT LOG")
	log.Println("	+ Xbox account required")
	log.Println("	+ Login command: xbox login")
	log.Println("	+ Spawn command: xbox join ip:port")
	log.Println("2. JOIN GAME")
	log.Println("	+ Must disable xbox-auth in server")
	log.Println("	+ Spawn command: join ip:port")
	log.Println("3. LIST")
	log.Println("	+ List all specters")
	log.Println("	+ Command: list")
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
				log.Println("[SPECTER] Login...")
				if err := xbl.InitializeToken(log); err != nil {
					log.Println("[SPECTER] Error: ", err)
				}
				log.Println("[SPECTER] Login success")
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
		}
	}
}
