package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupportsTLSHelper(t *testing.T) {
	assert.True(t, SupportsTLS("abba[mnop]qrst"))
	assert.False(t, SupportsTLS("abcd[bddb]xyyx"))
	assert.False(t, SupportsTLS("aaaa[qwer]tyui"))
	assert.True(t, SupportsTLS("ioxxoj[asdfgh]zxcvbn"))
}

func TestSupportsSSLHelper(t *testing.T) {
	assert.True(t, SupportsSSL("aba[bab]xyz"))
	assert.False(t, SupportsSSL("xyx[xyx]xyx"))
	assert.True(t, SupportsSSL("aaa[kek]eke"))
	assert.True(t, SupportsSSL("zazbz[bzb]cdb"))
}

func TestPartOne(t *testing.T) {
	assert.Equal(t, 118, PartOne("input.txt"))
}

func TestPartTwo(t *testing.T) {
	assert.Equal(t, 260, PartTwo("input.txt"))
}
