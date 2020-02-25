package parameter

import (
	"fmt"
	"net/url"
)

const (
	scheme = "amqp"
)

// Parameter represents the adapter parameter
type Parameter struct {
	Host     string
	Port     int8
	Username string
	Password string
	VHost    string
}

// New creates a new parameter instance
func New(params map[string]interface{}) (*Parameter, error) {
	return &Parameter{
		Host:     params["host"].(string),
		Port:     params["port"].(int8),
		Username: params["username"].(string),
		Password: params["password"].(string),
		VHost:    params["vhost"].(string),
	}, nil
	// FIXME: Check parameters
}

// URL returns the URL to connect
func (p *Parameter) URL() *url.URL {
	return &url.URL{
		Scheme: scheme,
		User:   url.UserPassword(p.Username, p.Password),
		Host:   fmt.Sprintf("%s:%d", p.Host, p.Port),
		Path:   p.VHost,
	}
}
