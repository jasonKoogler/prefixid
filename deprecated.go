// Package prefixid is deprecated and archived.
package prefixid // import "github.com/kromacorp/prefixid"

import (
	"fmt"
	"os"
)

func init() {
	fmt.Fprintln(os.Stderr, "WARNING: Package github.com/jasonKoogler/prefixid is archived and no longer maintained. Please use github.com/kromacorp/prefixid instead.")
}
