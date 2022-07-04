package specter

import (
	"github.com/phuongaz/easyspecter/xbl"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

func (s *SpecterXbox) Login(addr string) (specter *SpecterXbox, err error) {
	if err := xbl.InitializeToken(s.Log); err != nil {
		s.Log.Println("Error: ", err)
	}
	s.Log.Println("Login success")

	s.Log.Println("Join...")
	conn, err := minecraft.Dialer{
		ClientData:  login.ClientData{},
		TokenSource: xbl.TokenSrc,
	}.Dial("raknet", addr)

	s.Conn = conn
	if err != nil {
		return nil, err
	}
	s.Log.Println(conn.IdentityData().DisplayName + " joined")
	return s, nil
}
