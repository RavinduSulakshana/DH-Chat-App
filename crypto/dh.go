package crypto

import (
	"crypto/rand"
	"math/big"
)

var p, _ = new(big.Int).SetString("23", 10)
var g = big.NewInt(5)

func GenerateKeyPair() (*big.Int, *big.Int, error) {
	private, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	if err != nil {
		return nil, nil, err
	}
	private.Add(private, big.NewInt(2))
	public := new(big.Int).Exp(g, private, p)
	return private, public, nil
}
