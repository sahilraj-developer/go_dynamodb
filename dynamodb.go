package dynamodb

import (
	"fmt"
	"os"
)

// Config defines the minimum settings needed to initialize a DynamoDB client wrapper.
type Config struct {
	Region    string
	TableName string
}

// Client wraps the DynamoDB configuration for this package.
type Client struct {
	config Config
}

// NewClient creates a new client wrapper.
func NewClient(config Config) *Client {
	return &Client{config: config}
}

// Region returns the configured AWS region.
func (c *Client) Region() string {
	return c.config.Region
}

// Table returns the configured DynamoDB table name.
func (c *Client) Table() string {
	return c.config.TableName
}

// Config returns a copy of the client's current configuration.
func (c *Client) Config() Config {
	return c.config
}

// Validate checks that required settings are present.
func (c *Client) Validate() error {
	if c.config.Region == "" {
		return fmt.Errorf("region is required")
	}
	if c.config.TableName == "" {
		return fmt.Errorf("table name is required")
	}
	return nil
}

// NewClientFromEnv creates a client from AWS_REGION and DYNAMODB_TABLE_NAME.
func NewClientFromEnv() *Client {
	return NewClient(Config{
		Region:    os.Getenv("AWS_REGION"),
		TableName: os.Getenv("DYNAMODB_TABLE_NAME"),
	})
}
