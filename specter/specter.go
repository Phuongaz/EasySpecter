package specter

import (
	"log"

	"github.com/sandertv/gophertunnel/minecraft"
)

type Specter struct {
	Conn    *minecraft.Conn
	Log     *log.Logger
	Address string
}
type SpecterXbox struct {
	Specter
}

type SpecterNormal struct {
	Specter
	Name string
}

type SpecterInterface interface {
	Login(addr string) (specter *SpecterInterface, err error)
}
