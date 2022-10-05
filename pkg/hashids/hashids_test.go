package hashids

import (
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	asserts := assert.New(t)

	conf.SystemConfig.HashIDSalt = "test"
	Init()
	asserts.Equal(4, data.MinLength)
	asserts.Equal("test", data.Salt)

	encode, err := HashEncode([]int{1, 2})
	asserts.NoError(err)
	asserts.NotEmpty(encode)
}

func TestHashIDEncode(t *testing.T) {
	asserts := assert.New(t)

	encode := HashIDEncode(1, TypeUser)
	asserts.NotEmpty(encode)
}

func TestHashIDDecode(t *testing.T) {
	asserts := assert.New(t)

	{
		encode := HashIDEncode(1, TypeUser)
		asserts.NotEmpty(encode)

		decode, err := HashIDDecode(encode, TypeUser)
		asserts.NoError(err)
		asserts.Equal(uint(1), decode)
	}
	{
		decode, err := HashIDDecode("test", TypeUser)
		asserts.Error(err)
		asserts.Equal(uint(0), decode)
	}
}
