package token

import "time"

// this interface manages the token
type Maker interface {
	// username and expiry duration is encoded in to a token during creation
	CreateToken(username string, duration time.Duration) (string, error)

	// verify token will take the token check its validity, and return the paylod
	VerifyToken(token string) (*Payload, error)
}
