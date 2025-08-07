package models

import (
	"math/big"

	"github.com/gorilla/websocket"
)

type User struct {
	ID            string              `json:"id"`
	Username      string              `json:"username"`
	Conn          *websocket.Conn     `json:"-"`
	PublicKey     *big.Int            `json:"public_key,omitempty"`
	PrivateKey    *big.Int            `json:"-"`
	SharedSecrets map[string]*big.Int `json:"-"`
	AESKeys       map[string][]byte   `json:"-"`
}

func NewUser(id, username string, conn *websocket.Conn) *User {
	return &User{
		ID:            id,
		Username:      username,
		Conn:          conn,
		SharedSecrets: make(map[string]*big.Int),
		AESKeys:       make(map[string][]byte),
	}
}
