package duke_nukem

import (
	"crypto/rand"
	"encoding/base64"
)

func generatePassword() string {
	rb := make([]byte, 32)
	rand.Read(rb)
	return base64.URLEncoding.EncodeToString(rb)
}
