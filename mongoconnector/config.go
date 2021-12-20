package mongoconnector

import (
	"fmt"

	"github.com/gobuffalo/envy"
)

type Config struct {
	Username   string
	Password   string
	Database   string
	Collection string
	Host       string
}

// GetConnection returns connection string for Mongo.
func (c *Config) GetConnection() string {
	return fmt.Sprintf("mongodb://%s:%s@%s", c.Username, c.Password, c.Host)
}

// SetConfig tries to create a config from env variables.
func SetConfig() Config {
	c := Config{}

	user := envy.Get("MONGO_USER", "")
	password := envy.Get("MONGO_PASSWORD", "")
	host := envy.Get("MONGO_HOST", "")
	database := envy.Get("MONGO_DB", "")
	collection := envy.Get("MONGO_COLLECTION", "")

	if user == "" {
		panic("no credentials set")
	}

	c.Username = user
	c.Password = password
	c.Host = host
	c.Database = database
	c.Collection = collection

	return c
}
