package specter

import (
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

func (s *SpecterNormal) Login() (specter *SpecterNormal, err error) {
	conn, err := minecraft.Dialer{
		IdentityData: login.IdentityData{
			DisplayName: s.Name,
		},
		ErrorLog: s.Log,
	}.Dial("raknet", s.Address)

	if err != nil {
		return nil, err
	}
	s.Conn = conn
	s.Log.Println(conn.IdentityData().DisplayName + " joined")
	return s, nil
}
