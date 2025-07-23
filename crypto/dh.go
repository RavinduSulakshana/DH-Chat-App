package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

// use proper DH parameteres
var p, _ = new(big.Int).SetString("87A8E61DB4B6663CFFBBD19C651959998CEEF608660DD0F25D2CEED4435E3B00E00DF8F1D61957D4FAF7DF4561B2AA3016C3D91134096FAA3BF4296D830E9A7C209E0C6497517ABD5A8A9D306BCF67ED91F9E6725B4758C022E0B1EF4275BF7B6C5BFC11D45F90823DEBCDD0D9C40B4FC03DD4A7F9F61E09D6E3C9D46E077A1E29E40B5C9FF6A7F4F4C6A8A7F58BFEB6B96C7D7A8E2FB5F3E4EF8F9F1A4", 16)
var g = big.NewInt(2)

// generate DH private public key pair
func GenerateKeyPair() (*big.Int, *big.Int, error) {
	//generate private key
	private, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2)))
	if err != nil {
		return nil, nil, err
	}
	private.Add(private, big.NewInt(2))

	//calculate public key
	public := new(big.Int).Exp(g, private, p)

	return private, public, nil
}

// compute shared secret key
func SharedSecretKey(private, peerpublic *big.Int) *big.Int {
	return new(big.Int).Exp(peerpublic, private, p)
}

// derives 32 byte aes key from shared secret
func DerivesAesKey(sharedsecret *big.Int) []byte {
	hash := sha256.Sum256(sharedsecret.Bytes())
	return hash[:]
}
