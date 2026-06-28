# go_dynamodb

go_dynamodb is a lightweight Go package for configuring and validating DynamoDB-related settings in AWS-based applications.

It is especially useful for projects that combine DynamoDB with services such as S3 for file storage, upload tracking, and metadata management.

## Installation

```bash
go get github.com/sahilraj-developer/go_dynamodb
```

## Features

- Simple configuration for AWS region and DynamoDB table name
- Built-in validation for required settings
- Environment-based initialization using `AWS_REGION` and `DYNAMODB_TABLE_NAME`
- A clean foundation for building CRUD helpers, upload metadata flows, and DynamoDB-backed indexing features

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

You can also initialize the client directly from environment variables:

```go
client := dynamodb.NewClientFromEnv()
```

Supported variables:

- `AWS_REGION`
- `DYNAMODB_TABLE_NAME`

## Example Use Cases

- Storing S3 upload metadata in DynamoDB
- Tracking uploaded files and their status
- Building DynamoDB-backed indexes for assets or documents
- Simplifying AWS configuration in Go services

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome. Feel free to open issues or submit pull requests for improvements or new features.

