# prefixid

A Go library for managing prefixed IDs with type safety using generics.

## Features

- Type-safe ID prefixing and parsing using Go generics
- Thread-safe registry of entity types and their prefixes
- Support for custom ID types and prefixing strategies
- Easy initialization with predefined prefix maps
- Built-in prefixers for common ID types: string, int, UUID, ULID, KSUID

## Installation

```bash
go get github.com/jasonKoogler/prefixid
```

## Usage

### Basic usage with string IDs

```go
package main

import (
	"fmt"
	"github.com/jasonKoogler/prefixid"
)

func main() {
	// Create a new registry for string IDs
	registry := prefixid.NewRegistry[string]()

	// Register entity types with prefixes and the StringPrefixer
	stringPrefixer := prefixid.StringPrefixer{}
	registry.Register("user", "usr", stringPrefixer)
	registry.Register("product", "prod", stringPrefixer)

	// Create prefixed IDs
	userID, _ := registry.PrefixID("user", "12345")
	fmt.Println("Prefixed user ID:", userID) // usr_12345

	// Parse prefixed IDs
	parsedUserID, _ := registry.ParsePrefixedID("user", userID)
	fmt.Println("Parsed user ID:", parsedUserID) // 12345

	// Determine entity type from a prefixed ID
	entityType, rawID, ok := registry.MatchPrefix(userID)
	if ok {
		fmt.Printf("Entity type: %s, Raw ID: %s\n", entityType, rawID)
	}
}
```

### Using with integer IDs

```go
registry := prefixid.NewRegistry[int]()
intPrefixer := prefixid.IntPrefixer{}
registry.Register("customer", "cust", intPrefixer)

// Create a prefixed ID from an integer
customerID, _ := registry.PrefixID("customer", 42)
fmt.Println(customerID) // cust_42

// Parse back to an integer (type-safe)
parsedID, _ := registry.ParsePrefixedID("customer", customerID)
fmt.Printf("%d (type: %T)\n", parsedID, parsedID) // 42 (type: int)
```

### Using UUID prefixer

```go
import (
	"github.com/google/uuid"
	"github.com/jasonKoogler/prefixid"
)

// Create a registry for UUID IDs
registry := prefixid.NewRegistry[uuid.UUID]()
uuidPrefixer := prefixid.UUIDPrefixer{}
registry.Register("order", "ord", uuidPrefixer)

// Create a prefixed ID from a UUID
orderUUID := uuid.New()
orderID, _ := registry.PrefixID("order", orderUUID)
fmt.Println(orderID) // ord_6ba7b810-9dad-11d1-80b4-00c04fd430c8

// Parse back to a UUID (type-safe)
parsedUUID, _ := registry.ParsePrefixedID("order", orderID)
fmt.Printf("%s (type: %T)\n", parsedUUID, parsedUUID)
```

### Using ULID prefixer

```go
import (
	"github.com/oklog/ulid/v2"
	"github.com/jasonKoogler/prefixid"
)

// Create a registry for ULID IDs
registry := prefixid.NewRegistry[ulid.ULID]()
ulidPrefixer := prefixid.ULIDPrefixer{}
registry.Register("event", "evt", ulidPrefixer)

// Create a prefixed ID from a ULID
eventULID := ulid.MustNew(ulid.Now(), nil)
eventID, _ := registry.PrefixID("event", eventULID)
fmt.Println(eventID) // evt_01H2XEWE6NMWBGR9RZGWMSD4PQ

// Parse back to a ULID (type-safe)
parsedULID, _ := registry.ParsePrefixedID("event", eventID)
fmt.Printf("%s (type: %T)\n", parsedULID, parsedULID)
```

### Using KSUID prefixer

```go
import (
	"github.com/segmentio/ksuid"
	"github.com/jasonKoogler/prefixid"
)

// Create a registry for KSUID IDs
registry := prefixid.NewRegistry[ksuid.KSUID]()
ksuidPrefixer := prefixid.KSUIDPrefixer{}
registry.Register("transaction", "txn", ksuidPrefixer)

// Create a prefixed ID from a KSUID
txnKSUID := ksuid.New()
txnID, _ := registry.PrefixID("transaction", txnKSUID)
fmt.Println(txnID) // txn_1sYfwSUNHYvThjGfJH1tFdgYQTb

// Parse back to a KSUID (type-safe)
parsedKSUID, _ := registry.ParsePrefixedID("transaction", txnID)
fmt.Printf("%s (type: %T)\n", parsedKSUID, parsedKSUID)
```

### Using predefined prefix maps

```go
prefixMap := map[string]string{
    "user":    "usr",
    "product": "prod",
    "order":   "ord",
}

// Create a registry with predefined prefixes
registry := prefixid.NewRegistryWithPrefixes[string](prefixMap)

// Register prefixers for each entity type
stringPrefixer := prefixid.StringPrefixer{}
for entityType, prefix := range prefixMap {
    registry.Register(entityType, prefix, stringPrefixer)
}
```

## Creating custom prefixers

You can implement the `IDPrefixer` interface for any custom ID type:

```go
type IDPrefixer[T any] interface {
    // Attach attaches a prefix to an ID
    Attach(prefix string, id T) string
    
    // Detach detaches a prefix from a prefixed ID
    Detach(prefix string, prefixedID string) (string, bool)
    
    // Parse parses a string into an ID
    Parse(s string) (T, error)
}
```

Example of a custom prefixer:

```go
// MyCustomPrefixer for a custom ID type
type MyCustomPrefixer struct{}

// Attach adds a prefix to a custom ID
func (p MyCustomPrefixer) Attach(prefix string, id MyCustomID) string {
    return fmt.Sprintf("%s:%s", prefix, id.String())
}

// Detach removes the prefix from a prefixed ID string
func (p MyCustomPrefixer) Detach(prefix string, prefixedID string) (string, bool) {
    expectedPrefix := fmt.Sprintf("%s:", prefix)
    if strings.HasPrefix(prefixedID, expectedPrefix) {
        return strings.TrimPrefix(prefixedID, expectedPrefix), true
    }
    return "", false
}

// Parse converts a string to a custom ID type
func (p MyCustomPrefixer) Parse(s string) (MyCustomID, error) {
    // Implementation for parsing the string to your custom ID type
}
```

## Testing

A comprehensive test suite is included and can be run with:

```bash
make test
```

For test coverage:

```bash
make cover
```

The tests are located in the `/tests` directory and provide examples of how to use each feature of the library.

## License

MIT
