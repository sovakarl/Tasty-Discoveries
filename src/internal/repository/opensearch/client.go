package opensearch

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/opensearch-project/opensearch-go/v4"
)

var retryOnStatus []int = []int{502, 503, 504}
const maxRetries = 5
const skipVerify = true

type Config struct {
	Host     string
	Port     string
	Password string
}

func (c Config) address() []string {
	return []string{fmt.Sprintf("%s:%s", c.Host, c.Port)}
}

type Client struct {
	clinet *opensearch.Client
}

func NewClient(cnf Config) (*Client, error) {
	client, err := opensearch.NewClient(
		opensearch.Config{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify}},
			Addresses:     cnf.address(),
			Password:      cnf.Password,
			MaxRetries:    maxRetries,
			RetryOnStatus: retryOnStatus,
		},
	)

	if err != nil {
		return nil, err
	}
	result := &Client{clinet: client}
	return result, nil
}
