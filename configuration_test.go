package qovery

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestInitQoveryWithFile(t *testing.T) {
	genericTestFunc := func(q *Qovery) {
		assert.NotNil(t, q, "it should be not nil")
		assert.Equal(t, "master", q.GetBranchName(), "it should return master")
		assert.True(t, q.IsProduction(), "it should return true")
		assert.NotEmpty(t, q.GetDatabaseConfigurations(), "list should not be empty")
		assert.NotNil(t, q.GetDatabaseConfigurationByName("my-pql"), "my-pql database should exist")
		assert.Nil(t, q.GetDatabaseConfigurationByName("sql-server"), "sql-server database should not exist")
	}

	os.Setenv(EnvIsProduction, "true")
	os.Setenv(EnvBranchName, "master")

	t.Log("test using configuration file")
	filename := "./test_files/local_configuration.json"
	q, err := New(&filename)
	assert.Nil(t, err, "it should be nil")
	genericTestFunc(q)

	t.Log("test using configuration base 64 env")
	file, _ := ioutil.ReadFile("test_files/b64.txt")
	os.Setenv(EnvJsonB64, string(file))
	q, err = New(nil)
	assert.Nil(t, err, "it should be nil")
	genericTestFunc(q)
}
