package token

import "time"

// maker is interface for managing tokens
type Maker interface {
	// CreateToken create new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// Checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
