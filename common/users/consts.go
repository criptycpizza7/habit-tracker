package users

import "time"

const (
	HASH_COST = 10

	JWT_TOKEN_TTL = time.Hour * 24 * 7
)

var JWT_SECRET = []byte("secret") // TODO: грузить из .env
