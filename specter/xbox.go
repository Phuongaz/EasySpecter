package specter

import (
	"context"

	"github.com/phuongaz/easyspecter/xbl"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

func (s *SpecterXbox) Login() (specter *SpecterXbox, err error) {
	if err := xbl.InitializeToken(s.Log); err != nil {
		s.Log.Println("Error: ", err)
	}
	s.Log.Println("Login success")
	s.Log.Println("Join...")
	ctx := context.Background()

	conn, err := minecraft.Dialer{
		ClientData:  login.ClientData{},
		TokenSource: xbl.TokenSrc,
	}.DialContext(ctx, "raknet", s.Address)

	s.Conn = conn
	if err != nil {
		return nil, err
	}
	s.Log.Println(conn.IdentityData().DisplayName + " joined")
	return s, nil
}
