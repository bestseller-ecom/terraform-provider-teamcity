package teamcity

import (
	"errors"

	"github.com/bestseller-ecom/teamcity-sdk-go/teamcity"
	// "os"
)

type Config struct {
	User                string
	Password            string
	URL                 string
	Version             string
	Insecure            bool
	SkipCredsValidation bool
}

func (c *Config) Client() (interface{}, error) {
	// if c.User == "" {
	// 	c.User = os.Getenv("TEAMCITY_USERNAME")
	// }
	if c.User == "" {
		return nil, errors.New("Missing TeamCity user and TEAMCITY_USERNAME not defined")
	}

	// if c.Password == "" {
	// 	c.Password = os.Getenv("TEAMCITY_PASSWORD")
	// }
	if c.Password == "" {
		return nil, errors.New("Missing TeamCity password and TEAMCITY_PASSWORD not defined")
	}

	// if c.URL == "" {
	// 	c.URL = os.Getenv("TEAMCITY_URL")
	// }
	if c.URL == "" {
		return nil, errors.New("Missing TeamCity URL and TEAMCITY_URL not defined")
	}

	if c.Version == "" {
		return nil, errors.New("Missing TeamCity API Version and TEAMCITY_API_VERSION not defined")
	}

	client := teamcity.New(c.URL, c.User, c.Password, c.Version)

	if !c.SkipCredsValidation {
		err := c.ValidateCredentials(client)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// Validate credentials early and fail before we do any graph walking.
func (c *Config) ValidateCredentials(client *teamcity.Client) error {
	server, err := client.Server()
	if err != nil {
		return err
	}
	if server == nil {
		return errors.New("Received no reply from server")
	}
	return nil
}
