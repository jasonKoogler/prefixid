package prefixid_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jasonKoogler/prefixid"
	"github.com/oklog/ulid/v2"
	"github.com/segmentio/ksuid"
)

func TestIntegration_AllPrefixers(t *testing.T) {
	// Create a registry for string IDs
	strRegistry := prefixid.NewRegistry[string]()
	strRegistry.Register("user", "usr", prefixid.StringPrefixer{})

	// Create a registry for int IDs
	intRegistry := prefixid.NewRegistry[int]()
	intRegistry.Register("product", "prd", prefixid.IntPrefixer{})

	// Create a registry for UUID IDs
	uuidRegistry := prefixid.NewRegistry[uuid.UUID]()
	uuidRegistry.Register("order", "ord", prefixid.UUIDPrefixer{})

	// Create a registry for ULID IDs
	ulidRegistry := prefixid.NewRegistry[ulid.ULID]()
	ulidRegistry.Register("session", "ses", prefixid.ULIDPrefixer{})

	// Create a registry for KSUID IDs
	ksuidRegistry := prefixid.NewRegistry[ksuid.KSUID]()
	ksuidRegistry.Register("transaction", "txn", prefixid.KSUIDPrefixer{})

	// Test string IDs
	userId := "abc123"
	userPrefixedID, err := strRegistry.PrefixID("user", userId)
	if err != nil {
		t.Errorf("Failed to prefix user ID: %v", err)
	}

	if userPrefixedID != "usr_abc123" {
		t.Errorf("Expected 'usr_abc123', got %s", userPrefixedID)
	}

	parsedUserID, err := strRegistry.ParsePrefixedID("user", userPrefixedID)
	if err != nil {
		t.Errorf("Failed to parse prefixed user ID: %v", err)
	}

	if parsedUserID != userId {
		t.Errorf("Expected '%s', got %s", userId, parsedUserID)
	}

	// Test int IDs
	productId := 42
	productPrefixedID, err := intRegistry.PrefixID("product", productId)
	if err != nil {
		t.Errorf("Failed to prefix product ID: %v", err)
	}

	if productPrefixedID != "prd_42" {
		t.Errorf("Expected 'prd_42', got %s", productPrefixedID)
	}

	parsedProductID, err := intRegistry.ParsePrefixedID("product", productPrefixedID)
	if err != nil {
		t.Errorf("Failed to parse prefixed product ID: %v", err)
	}

	if parsedProductID != productId {
		t.Errorf("Expected %d, got %d", productId, parsedProductID)
	}

	// Test UUID IDs
	orderId := uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")
	orderPrefixedID, err := uuidRegistry.PrefixID("order", orderId)
	if err != nil {
		t.Errorf("Failed to prefix order ID: %v", err)
	}

	if orderPrefixedID != "ord_f47ac10b-58cc-0372-8567-0e02b2c3d479" {
		t.Errorf("Expected 'ord_f47ac10b-58cc-0372-8567-0e02b2c3d479', got %s", orderPrefixedID)
	}

	parsedOrderID, err := uuidRegistry.ParsePrefixedID("order", orderPrefixedID)
	if err != nil {
		t.Errorf("Failed to parse prefixed order ID: %v", err)
	}

	if parsedOrderID != orderId {
		t.Errorf("Expected %s, got %s", orderId, parsedOrderID)
	}

	// Test ULID IDs
	sessionId, _ := ulid.Parse("01F8MECHZX3TBDSZ9PT3RV4ZMH")
	sessionPrefixedID, err := ulidRegistry.PrefixID("session", sessionId)
	if err != nil {
		t.Errorf("Failed to prefix session ID: %v", err)
	}

	if sessionPrefixedID != "ses_01F8MECHZX3TBDSZ9PT3RV4ZMH" {
		t.Errorf("Expected 'ses_01F8MECHZX3TBDSZ9PT3RV4ZMH', got %s", sessionPrefixedID)
	}

	parsedSessionID, err := ulidRegistry.ParsePrefixedID("session", sessionPrefixedID)
	if err != nil {
		t.Errorf("Failed to parse prefixed session ID: %v", err)
	}

	if parsedSessionID != sessionId {
		t.Errorf("Expected %s, got %s", sessionId, parsedSessionID)
	}

	// Test KSUID IDs
	transactionId, _ := ksuid.Parse("0ujtsYcgvSTl8PAuAdqWYSMnLOv")
	transactionPrefixedID, err := ksuidRegistry.PrefixID("transaction", transactionId)
	if err != nil {
		t.Errorf("Failed to prefix transaction ID: %v", err)
	}

	if transactionPrefixedID != "txn_0ujtsYcgvSTl8PAuAdqWYSMnLOv" {
		t.Errorf("Expected 'txn_0ujtsYcgvSTl8PAuAdqWYSMnLOv', got %s", transactionPrefixedID)
	}

	parsedTransactionID, err := ksuidRegistry.ParsePrefixedID("transaction", transactionPrefixedID)
	if err != nil {
		t.Errorf("Failed to parse prefixed transaction ID: %v", err)
	}

	if parsedTransactionID != transactionId {
		t.Errorf("Expected %s, got %s", transactionId, parsedTransactionID)
	}
}

func TestIntegration_MatchPrefix(t *testing.T) {
	// Create a mixed registry for string IDs
	registry := prefixid.NewRegistry[string]()
	registry.Register("user", "usr", prefixid.StringPrefixer{})
	registry.Register("post", "pst", prefixid.StringPrefixer{})
	registry.Register("comment", "cmt", prefixid.StringPrefixer{})

	// Test matching different prefixes
	testCases := []struct {
		prefixedID     string
		expectedEntity string
		expectedID     string
		expectMatch    bool
	}{
		{"usr_123", "user", "123", true},
		{"pst_abc", "post", "abc", true},
		{"cmt_xyz", "comment", "xyz", true},
		{"inv_456", "", "", false},
		{"usr123", "", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.prefixedID, func(t *testing.T) {
			entityType, rawID, ok := registry.MatchPrefix(tc.prefixedID)

			if ok != tc.expectMatch {
				t.Errorf("Expected match=%v, got %v", tc.expectMatch, ok)
			}

			if !tc.expectMatch {
				return
			}

			if entityType != tc.expectedEntity {
				t.Errorf("Expected entity type %s, got %s", tc.expectedEntity, entityType)
			}

			if rawID != tc.expectedID {
				t.Errorf("Expected raw ID %s, got %s", tc.expectedID, rawID)
			}
		})
	}
}
