package triggy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/julienbreux/triggy/internal/triggy"
)

func TestYamlUnmarshal(t *testing.T) {
	var ymlParameters = make(map[string]interface{})
	ymlParameters["host"] = "localhost"
	ymlParameters["port"] = 5672
	ymlParameters["username"] = "guest"
	ymlParameters["password"] = "guest"
	ymlParameters["vhost"] = "/"
	const (
		ymlProvider = `
adapter: rabbitmq
parameters:
  host: localhost
  port: 5672
  username: guest
  password: guest
  vhost: /`
		ymlProviderBadParameters = `
  adapter: rabbitmq
  parameters: cool`
		ymlProviderBadAdapter = `adapter: fake`
	)
	tests := []struct {
		content  string
		value    interface{}
		expected interface{}
		hasErr   bool
	}{
		{
			ymlProvider,
			&triggy.Provider{},
			&triggy.Provider{Adapter: "rabbitmq", Parameters: ymlParameters},
			false,
		},
		{
			ymlProviderBadParameters,
			&triggy.Provider{},
			&triggy.Provider{},
			true,
		},
		{
			ymlProviderBadAdapter,
			&triggy.Provider{},
			&triggy.Provider{},
			true,
		},
	}

	for _, test := range tests {
		err := yaml.Unmarshal([]byte(test.content), test.value)
		if !test.hasErr {
			assert.NoError(t, err)
		}
		assert.Equal(t, test.expected, test.value)
	}
}
