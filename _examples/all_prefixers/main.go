package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/jasonKoogler/prefixid"
	"github.com/oklog/ulid/v2"
	"github.com/segmentio/ksuid"
)

func main() {
	// Example with String IDs
	fmt.Println("=== String IDs ===")
	stringRegistry := prefixid.NewRegistry[string]()
	stringPrefixer := prefixid.StringPrefixer{}
	stringRegistry.Register("user", "usr", stringPrefixer)

	userID, _ := stringRegistry.PrefixID("user", "abc123")
	fmt.Println("Prefixed user ID:", userID)

	parsed, _ := stringRegistry.ParsePrefixedID("user", userID)
	fmt.Printf("Parsed: %s (type: %T)\n", parsed, parsed)

	// Example with Int IDs
	fmt.Println("\n=== Int IDs ===")
	intRegistry := prefixid.NewRegistry[int]()
	intPrefixer := prefixid.IntPrefixer{}
	intRegistry.Register("product", "prod", intPrefixer)

	productID, _ := intRegistry.PrefixID("product", 42)
	fmt.Println("Prefixed product ID:", productID)

	parsedInt, _ := intRegistry.ParsePrefixedID("product", productID)
	fmt.Printf("Parsed: %d (type: %T)\n", parsedInt, parsedInt)

	// Example with UUID IDs
	fmt.Println("\n=== UUID IDs ===")
	uuidRegistry := prefixid.NewRegistry[uuid.UUID]()
	uuidPrefixer := prefixid.UUIDPrefixer{}
	uuidRegistry.Register("order", "ord", uuidPrefixer)

	orderUUID := uuid.New()
	orderID, _ := uuidRegistry.PrefixID("order", orderUUID)
	fmt.Println("Original UUID:", orderUUID)
	fmt.Println("Prefixed order ID:", orderID)

	parsedUUID, _ := uuidRegistry.ParsePrefixedID("order", orderID)
	fmt.Printf("Parsed: %s (type: %T)\n", parsedUUID, parsedUUID)

	// Example with ULID IDs
	fmt.Println("\n=== ULID IDs ===")
	ulidRegistry := prefixid.NewRegistry[ulid.ULID]()
	ulidPrefixer := prefixid.ULIDPrefixer{}
	ulidRegistry.Register("event", "evt", ulidPrefixer)

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	eventULID := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
	eventID, _ := ulidRegistry.PrefixID("event", eventULID)
	fmt.Println("Original ULID:", eventULID)
	fmt.Println("Prefixed event ID:", eventID)

	parsedULID, _ := ulidRegistry.ParsePrefixedID("event", eventID)
	fmt.Printf("Parsed: %s (type: %T)\n", parsedULID, parsedULID)

	// Example with KSUID IDs
	fmt.Println("\n=== KSUID IDs ===")
	ksuidRegistry := prefixid.NewRegistry[ksuid.KSUID]()
	ksuidPrefixer := prefixid.KSUIDPrefixer{}
	ksuidRegistry.Register("transaction", "txn", ksuidPrefixer)

	txnKSUID := ksuid.New()
	txnID, _ := ksuidRegistry.PrefixID("transaction", txnKSUID)
	fmt.Println("Original KSUID:", txnKSUID)
	fmt.Println("Prefixed transaction ID:", txnID)

	parsedKSUID, _ := ksuidRegistry.ParsePrefixedID("transaction", txnID)
	fmt.Printf("Parsed: %s (type: %T)\n", parsedKSUID, parsedKSUID)
}
