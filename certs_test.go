package certs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRemoteCertificate(t *testing.T) {
	cert, err := ParseRemoteCertificate("google.com")

	assert.Nil(t, err)
	assert.Equal(t, true, strings.Contains(cert.CommonName, "google.com"), "should be true")
}
