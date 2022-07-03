package specter

import (
	"github.com/phuongaz/easyspecter/xbl"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

func (s *SpecterXbox) Login(addr string) (specter *SpecterXbox, err error) {
	s.Log.Println("SPECTER: Login...")
	if err := xbl.InitializeToken(s.Log); err != nil {
		s.Log.Println("SPECTER: Error: ", err)
	}
	s.Log.Println("SPECTER: Login success")

	s.Log.Println("SPECTER: Join...")
	conn, err := minecraft.Dialer{
		ClientData:  login.ClientData{},
		TokenSource: xbl.TokenSrc,
	}.Dial("raknet", addr)

	s.Conn = conn
	if err != nil {
		s.Log.Println("SPECTER: Error: ", err)
		return nil, err
	}
	s.Log.Println("SPECTER: " + conn.IdentityData().DisplayName + " joined")
	return s, nil
}
