package main

import (
	"fmt"

	"github.com/jasonKoogler/prefixid"
)

func main() {
	// Create a map of entity types to prefixes
	prefixMap := map[string]string{
		"user":    "usr",
		"product": "prod",
		"order":   "ord",
	}

	// Create a new registry with predefined prefixes
	registry := prefixid.NewRegistryWithPrefixes[string](prefixMap)

	// Register prefixers for each entity type
	stringPrefixer := prefixid.StringPrefixer{}
	registry.Register("user", prefixMap["user"], stringPrefixer)
	registry.Register("product", prefixMap["product"], stringPrefixer)
	registry.Register("order", prefixMap["order"], stringPrefixer)

	// Get a list of all registered entity types
	entityTypes := registry.GetEntityTypes()
	fmt.Println("Registered entity types:", entityTypes)

	// Create prefixed IDs
	userID, _ := registry.PrefixID("user", "12345")
	productID, _ := registry.PrefixID("product", "abcde")
	orderID, _ := registry.PrefixID("order", "xyz789")

	fmt.Println("Prefixed user ID:", userID)
	fmt.Println("Prefixed product ID:", productID)
	fmt.Println("Prefixed order ID:", orderID)

	// Parse prefixed IDs
	parsedUserID, _ := registry.ParsePrefixedID("user", userID)
	fmt.Println("Parsed user ID:", parsedUserID)

	// Match a prefix to determine entity type
	entityType, rawID, ok := registry.MatchPrefix(productID)
	if ok {
		fmt.Printf("Matched prefix: entity type = %s, raw ID = %s\n", entityType, rawID)
	}
}
