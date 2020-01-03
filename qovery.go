package qovery

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

const (
	EnvJsonB64                   = "QOVERY_JSON_B64"
	EnvIsProduction              = "QOVERY_IS_PRODUCTION"
	EnvBranchName                = "QOVERY_BRANCH_NAME"
	DefaultConfigurationFilename = ".qovery/local_configuration.json"
)

func New(configurationFilename *string) (*Qovery, error) {
	q := Qovery{}
	q.configuration = q.getConfigurationFromEnv(EnvJsonB64)
	if q.configuration == nil {
		filename := DefaultConfigurationFilename
		if configurationFilename != nil {
			filename = *configurationFilename
		}
		var err error
		q.configuration, err = q.getConfigurationFromFile(filename)
		return &q, err
	}
	return &q, nil
}

type Qovery struct {
	configuration *Configuration
}

func (q *Qovery) getConfigurationFromEnv(environmentVariable string) *Configuration {
	if environmentVariable == "" {
		return nil
	}
	b64JSon := os.Getenv(environmentVariable)
	if b64JSon == "" {
		return nil
	}
	jsonBytes, err := base64.URLEncoding.DecodeString(b64JSon)
	if err != nil {
		return nil
	}
	var conf Configuration
	if err := json.Unmarshal(jsonBytes, &conf); err != nil {
		return nil
	}
	return &conf
}

func (q *Qovery) getConfigurationFromFile(filename string) (*Configuration, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if err != os.ErrNotExist {
			return nil, err
		}
		return nil, nil
	}
	var conf Configuration
	if err := json.Unmarshal(file, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func (q *Qovery) GetConfiguration() *Configuration {
	return q.configuration
}

func (q *Qovery) IsProduction() bool {
	isProdStr := os.Getenv(EnvIsProduction)
	return strings.ToLower(isProdStr) == "true"
}

func (q *Qovery) GetBranchName() string {
	return os.Getenv(EnvBranchName)
}

func (q *Qovery) GetDatabaseConfigurations() []DatabaseConfiguration {
	if q.configuration == nil || q.configuration.Databases == nil {
		return []DatabaseConfiguration{}
	}
	return q.configuration.Databases
}

func (q *Qovery) GetDatabaseConfigurationByName(name string) *DatabaseConfiguration {
	if q.configuration == nil || q.configuration.Databases == nil {
		return nil
	}
	for _, db := range q.configuration.Databases {
		if db.Name == name {
			return &db
		}
	}
	return nil
}
