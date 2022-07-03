package specter

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

var specters = map[string]*Specter{}

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

func (s *Specter) Chat(message string) {
	s.Conn.WritePacket(&packet.Text{TextType: packet.TextTypeChat, Message: message})
}

func (s *Specter) Move(x, y, z float32) {
	s.Conn.WritePacket(&packet.CorrectPlayerMovePrediction{
		Position: mgl32.Vec3{x, y, z},
		OnGround: true,
		Tick:     uint64(20),
	})
}

func GetSpecters() map[string]*Specter {
	return specters
}

func AddSpecter(s *Specter) {
	specters[s.Conn.IdentityData().DisplayName] = s
}
