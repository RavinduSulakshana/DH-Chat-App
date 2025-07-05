package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

type DHKeyPair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func GenerateDHKeyPair() (*DHKeyPair, error) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return &DHKeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}, nil
}

// compute shared secret key using local private key and remote private key
func ComputeSharedSecret(privateKey *ecdsa.PrivateKey, remotePubKey *ecdsa.PublicKey) ([]byte, error) {
	x, _ := elliptic.P256().ScalarMult(remotePubKey.X, remotePubKey.Y, privateKey.D.Bytes())
	sharedSecret := sha256.Sum256(x.Bytes())
	return sharedSecret[:], nil
}

// Encode publicKey to Bytes for transmission
func EncodePublicKey(pubKey *ecdsa.PublicKey) []bytes {
	return elliptic.Marshal(elliptic.P256(), pubKey.X, pubKey.Y)
}

// Decode bytes to public key
func DecodePublicKey(data []bytes) *ecdsa.PublicKey {
	x, y := elliptic.Unmarshal(elliptic.P256(), data)
	return &ecdsa.PublicKey{Curve: elliptic.P224(), X: x, Y: y}

}
