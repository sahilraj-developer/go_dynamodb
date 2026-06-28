package dynamodb

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(Config{Region: "us-east-1", TableName: "uploads"})

	if client == nil {
		t.Fatal("expected client to be initialized")
	}

	if got := client.Region(); got != "us-east-1" {
		t.Fatalf("expected region us-east-1, got %q", got)
	}

	if got := client.Table(); got != "uploads" {
		t.Fatalf("expected table uploads, got %q", got)
	}
}

func TestValidateRequiresRegionAndTable(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{name: "valid", config: Config{Region: "us-east-1", TableName: "uploads"}, wantErr: false},
		{name: "missing region", config: Config{TableName: "uploads"}, wantErr: true},
		{name: "missing table", config: Config{Region: "us-east-1"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.config)
			err := client.Validate()
			if tt.wantErr && err == nil {
				t.Fatal("expected validation error")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}

func TestConfigReturnsCopy(t *testing.T) {
	client := NewClient(Config{Region: "us-east-1", TableName: "uploads"})

	cfg := client.Config()
	cfg.TableName = "changed"

	if got := client.Table(); got != "uploads" {
		t.Fatalf("expected immutable table value, got %q", got)
	}
}

func TestNewClientFromEnv(t *testing.T) {
	t.Setenv("AWS_REGION", "eu-west-1")
	t.Setenv("DYNAMODB_TABLE_NAME", "metadata")

	client := NewClientFromEnv()
	if got := client.Region(); got != "eu-west-1" {
		t.Fatalf("expected region from env, got %q", got)
	}
	if got := client.Table(); got != "metadata" {
		t.Fatalf("expected table from env, got %q", got)
	}
}

func TestNewClientFromEnvUsesExistingConfig(t *testing.T) {
	os.Unsetenv("AWS_REGION")
	os.Unsetenv("DYNAMODB_TABLE_NAME")

	client := NewClientFromEnv()
	if got := client.Region(); got != "" {
		t.Fatalf("expected empty region when env is unset, got %q", got)
	}
	if got := client.Table(); got != "" {
		t.Fatalf("expected empty table when env is unset, got %q", got)
	}
}
