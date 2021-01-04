package certs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRemoteCertificate(t *testing.T) {
	cert, err := ParseRemoteCertificate("google.com:443")

	assert.Nil(t, err)
	assert.Equal(t, true, strings.Contains(cert.CommonName, "google.com"), "should be true")
}
