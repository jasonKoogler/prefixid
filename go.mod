// Deprecated: This module is archived and no longer maintained.
// Please use github.com/kromacorp/prefixid instead.
module github.com/jasonKoogler/prefixid

go 1.24.2

require (
	github.com/google/uuid v1.6.0
	github.com/oklog/ulid/v2 v2.1.0
	github.com/segmentio/ksuid v1.0.4
)

// Indicate that all versions are deprecated
retract (
	v0.0.0-00010101000000-000000000000/latest  // All versions retracted: Package is archived
)
