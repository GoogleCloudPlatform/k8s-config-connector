# Storage Abstraction Layer

This package provides an abstract storage layer for the MultiCluster Lease system, allowing different storage backends to be plugged in for storing lease data.

## Interface

The `Storage` interface defines the operations needed for lease management:

```go
type Storage interface {
    ReadLease(ctx context.Context, key string) (*LeaseObject, error)
    WriteLease(ctx context.Context, key string, data *LeaseData, generation int64) error
    CreateLease(ctx context.Context, key string, data *LeaseData) error
    DeleteLease(ctx context.Context, key string) error
}
```

## Implementations

### 1. GCS Storage (`gcs.go`)

Implements the storage interface using Google Cloud Storage. This is the production-ready implementation for cloud deployments.

**Features:**
- Uses GCS object generation for optimistic locking
- Stores lease data as JSON objects
- Supports conditional operations to prevent race conditions
- Production-ready with proper error handling

**Usage:**
```go
import (
    "cloud.google.com/go/storage"
    "github.com/GoogleCloudPlatform/k8s-config-connector/experiments/multiclusterlease/pkg/storage"
)

client, _ := storage.NewClient(ctx)
gcsStorage := storage.NewGCSStorage(client, "my-bucket")
```

### 2. Memory Storage (`memory.go`)

In-memory implementation useful for testing and development.

**Features:**
- Thread-safe using mutex
- Simulates optimistic locking with generation counters
- No external dependencies
- Perfect for unit tests and local development

**Usage:**
```go
memStorage := storage.NewMemoryStorage()
```

## Implementing Your Own Storage Backend

To implement a custom storage backend:

1. Implement the `Storage` interface
2. Handle the specific error types:
   - Return `ErrNotFound` when a lease doesn't exist
   - Return `ErrAlreadyExists` when creating a lease that already exists
   - Return `ErrConditionalUpdateFailed` when optimistic locking fails
3. Ensure thread-safety if your backend doesn't provide it
4. Implement proper optimistic locking using whatever mechanism your storage provides

### Example: Redis Implementation

```go
type RedisStorage struct {
    client *redis.Client
}

func (r *RedisStorage) ReadLease(ctx context.Context, key string) (*LeaseObject, error) {
    // Use Redis HGETALL to read lease data and version
    // Parse JSON and return LeaseObject
}

func (r *RedisStorage) WriteLease(ctx context.Context, key string, data *LeaseData, generation int64) error {
    // Use Redis Lua script for atomic compare-and-swap based on generation
}

// ... implement other methods
```

## Error Handling

The storage layer defines standard errors that all implementations should use:

- `ErrNotFound`: Lease doesn't exist
- `ErrAlreadyExists`: Lease already exists during creation
- `ErrConditionalUpdateFailed`: Optimistic locking failed

Use the helper functions to check error types:
- `IsNotFound(err)`
- `IsAlreadyExists(err)`
- `IsConditionalUpdateFailed(err)`

## Data Model

### LeaseData

The actual lease information stored:

```go
type LeaseData struct {
    HolderIdentity   string    `json:"holderIdentity"`
    RenewTime        time.Time `json:"renewTime"`
    LeaseTransitions int32     `json:"leaseTransitions"`
}
```

### LeaseObject

Wraps lease data with metadata for optimistic locking:

```go
type LeaseObject struct {
    Data       LeaseData
    Generation int64 // Used for optimistic locking
}
```

## Thread Safety

All storage implementations must be thread-safe as they may be called concurrently by multiple goroutines within the controller. 