package remote

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDataFromGitHub(t *testing.T) {
	test := assert.New(t)
	Endpoint = "https://raw.fastgit.org/lezi-wiki/lezi-api/master/data.json"
	data, err := GetDataFromGitHub()
	test.NoError(err)
	test.NotNil(data)
	_, err = json.Marshal(data)
	test.NoError(err)
}
