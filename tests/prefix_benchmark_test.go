package prefixid_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jasonKoogler/prefixid"
	"github.com/oklog/ulid/v2"
	"github.com/segmentio/ksuid"
)

func setupStringRegistry() *prefixid.Registry[string] {
	registry := prefixid.NewRegistry[string]()
	registry.Register("user", "usr", prefixid.StringPrefixer{})
	registry.Register("post", "pst", prefixid.StringPrefixer{})
	registry.Register("comment", "cmt", prefixid.StringPrefixer{})
	return registry
}

func setupUUIDRegistry() *prefixid.Registry[uuid.UUID] {
	registry := prefixid.NewRegistry[uuid.UUID]()
	registry.Register("order", "ord", prefixid.UUIDPrefixer{})
	return registry
}

func setupULIDRegistry() *prefixid.Registry[ulid.ULID] {
	registry := prefixid.NewRegistry[ulid.ULID]()
	registry.Register("session", "ses", prefixid.ULIDPrefixer{})
	return registry
}

func setupKSUIDRegistry() *prefixid.Registry[ksuid.KSUID] {
	registry := prefixid.NewRegistry[ksuid.KSUID]()
	registry.Register("transaction", "txn", prefixid.KSUIDPrefixer{})
	return registry
}

func BenchmarkRegister(b *testing.B) {
	registry := prefixid.NewRegistry[string]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		registry.Register("user", "usr", prefixid.StringPrefixer{})
	}
}

func BenchmarkStringPrefixer_Prefix(b *testing.B) {
	prefixer := prefixid.StringPrefixer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefixer.Attach("usr", "123")
	}
}

func BenchmarkStringPrefixer_Unprefix(b *testing.B) {
	prefixer := prefixid.StringPrefixer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefixer.Detach("usr", "usr_123")
	}
}

func BenchmarkUUIDPrefixer_Prefix(b *testing.B) {
	prefixer := prefixid.UUIDPrefixer{}
	id := uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefixer.Attach("ord", id)
	}
}

func BenchmarkUUIDPrefixer_Unprefix(b *testing.B) {
	prefixer := prefixid.UUIDPrefixer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefixer.Detach("ord", "ord_f47ac10b-58cc-0372-8567-0e02b2c3d479")
	}
}

func BenchmarkPrefixID_String(b *testing.B) {
	registry := setupStringRegistry()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = registry.PrefixID("user", "123")
	}
}

func BenchmarkPrefixID_UUID(b *testing.B) {
	registry := setupUUIDRegistry()
	id := uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = registry.PrefixID("order", id)
	}
}

func BenchmarkParsePrefixedID_String(b *testing.B) {
	registry := setupStringRegistry()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = registry.ParsePrefixedID("user", "usr_123")
	}
}

func BenchmarkParsePrefixedID_UUID(b *testing.B) {
	registry := setupUUIDRegistry()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = registry.ParsePrefixedID("order", "ord_f47ac10b-58cc-0372-8567-0e02b2c3d479")
	}
}

func BenchmarkMatchPrefix_String(b *testing.B) {
	registry := setupStringRegistry()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = registry.MatchPrefix("usr_123")
	}
}

func BenchmarkMatchPrefix_LargeRegistry(b *testing.B) {
	// Create a registry with many entity types
	registry := prefixid.NewRegistry[string]()

	for i := 0; i < 100; i++ {
		registry.Register(
			"entity"+string(rune(i)),
			"pfx"+string(rune(i)),
			prefixid.StringPrefixer{},
		)
	}

	registry.Register("target", "tgt", prefixid.StringPrefixer{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Worst case scenario - the target prefix is the last one
		_, _, _ = registry.MatchPrefix("tgt_123")
	}
}
