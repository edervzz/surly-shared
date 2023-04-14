package function

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}
