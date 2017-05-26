package cypher

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitKeyCypher(t *testing.T) {
	InitKeyCypher()
	require.Len(t,Keycypher,16)
	assert.NotNil(t,Keycypher)
}

func TestEncryptDecriptString(t *testing.T) {
	InitKeyCypher()
	srcString := "basic origin test"
	cypString := EncryptString(srcString)

	require.NotEqual(t,cypString,srcString)
	assert.NotEmpty(t,cypString)

	dstString := DecryptString(cypString)

	assert.Equal(t,dstString,srcString)
}