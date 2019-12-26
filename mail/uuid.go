package misc

import (
	"encoding/hex"

	"github.com/google/uuid"
)

// UUID will generate a UUID
func UUID() string {
	uu, err := uuid.NewRandom()
	if err != nil {
		panic("UUID is fucked")
	}

	return hex.EncodeToString(uu[:])
}
