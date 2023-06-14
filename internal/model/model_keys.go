package model

import "time"

type KeyPair struct {
	ID         int
	PublicKey  []byte
	PrivateKey []byte
	CreatedAt  time.Time

	// Additional fields can be added as needed
}
