package main

import (
	"fmt"

	"github.com/jasonKoogler/prefixid"
)

func main() {
	// Create a new registry
	registry := prefixid.NewRegistry[int]()

	// Register entity types with prefixes and the IntPrefixer
	intPrefixer := prefixid.IntPrefixer{}
	registry.Register("customer", "cust", intPrefixer)
	registry.Register("invoice", "inv", intPrefixer)

	// Create prefixed IDs from integer IDs
	customerID, _ := registry.PrefixID("customer", 42)
	invoiceID, _ := registry.PrefixID("invoice", 1001)

	fmt.Println("Prefixed customer ID:", customerID) // cust_42
	fmt.Println("Prefixed invoice ID:", invoiceID)   // inv_1001

	// Parse prefixed IDs back to integers
	parsedCustomerID, _ := registry.ParsePrefixedID("customer", customerID)
	fmt.Printf("Parsed customer ID: %d (type: %T)\n", parsedCustomerID, parsedCustomerID)

	// Match a prefix to determine entity type
	entityType, rawID, ok := registry.MatchPrefix(invoiceID)
	if ok {
		fmt.Printf("Matched prefix: entity type = %s, raw ID = %s\n", entityType, rawID)

		// Parse the raw ID back to an integer
		parsedID, _ := intPrefixer.Parse(rawID)
		fmt.Printf("Parsed ID from matched prefix: %d\n", parsedID)
	}
}
