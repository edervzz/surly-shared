package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Pass Tests */
func Test_NewToken_OK(t *testing.T) {
	// given
	id := "12345"
	exp := "10"
	iss := "issuerTest"
	sub := "subjectTest"
	aud := "audienceTest"
	key := "Pa$$w0rd"
	// when
	actual := NewToken(id, exp, iss, sub, aud, key)
	// then
	assert.NotNil(t, actual)
}
