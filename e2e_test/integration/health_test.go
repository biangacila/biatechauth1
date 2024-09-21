package integration

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

type HealthSuite struct {
	suite.Suite
}

func TestHealthSuite(t *testing.T) {
	suite.Run(t, new(HealthSuite))
}

func (s *HealthSuite) TestHealthEndpointLocalhostFail() {
	var server = "localhost"
	var port = "8080"
	var prefix = "backend-biatechauth1/api"
	url := fmt.Sprintf("http://%s:%s/%s/status1", server, port, prefix)

	c := http.Client{}
	resp, err := c.Get(url)
	s.NoError(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	s.Equal(http.StatusOK, resp.StatusCode)
}
func (s *HealthSuite) TestHealthEndpointLocalhostPass() {
	var server = "localhost"
	var port = "8080"
	var prefix = "backend-biatechauth1/api"
	url := fmt.Sprintf("http://%s:%s/%s/status", server, port, prefix)

	c := http.Client{}
	resp, err := c.Get(url)
	s.NoError(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	s.Equal(http.StatusOK, resp.StatusCode)
}
