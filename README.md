# go_dynamodb

A lightweight Go package for working with DynamoDB configuration, validation, and AWS-backed metadata workflows.

This package is designed for applications that pair DynamoDB with services such as S3 for file storage, upload tracking, and metadata indexing.

## Installation

```bash
go get github.com/sahilraj-developer/go_dynamodb
```

## Module

```go
module github.com/sahilraj-developer/go_dynamodb
```

## Features

- Simple configuration wrapper for AWS region and table name
- Built-in validation for required settings
- Environment-based initialization with `AWS_REGION` and `DYNAMODB_TABLE_NAME`
- Easy to extend for CRUD operations, upload metadata, and indexing workflows

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/sahilraj-developer/go_dynamodb"
)

func main() {
    client := dynamodb.NewClient(dynamodb.Config{
        Region:    "us-east-1",
        TableName: "uploads",
    })

    if err := client.Validate(); err != nil {
        panic(err)
    }

    fmt.Printf("Region: %s\nTable: %s\n", client.Region(), client.Table())
}
```

## Environment Configuration

You can also initialize the client from environment variables:

```go
client := dynamodb.NewClientFromEnv()
```

Supported environment variables:

- `AWS_REGION`
- `DYNAMODB_TABLE_NAME`

## Example Use Cases

- S3 upload metadata storage
- File tracking and record management
- DynamoDB-backed indexing for uploaded assets
- AWS application configuration helpers

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome. Feel free to open issues or submit pull requests for new features and improvements.

