package grid

import (
	"fmt"
	"strings"
)

// Format of all GRIDs
const Format = "grid:namespace:owner:resource_type:resource_id"

// GRID (Global Resource ID) objects are used to uniquely identify resources.
type GRID struct {
	Namespace    string
	Owner        string
	ResourceType string
	ResourceID   string
}

// New returns a new GRID object.
func New(namespace, owner, resourceType, resourceID string) GRID {
	return GRID{
		Namespace:    namespace,
		Owner:        owner,
		ResourceType: resourceType,
		ResourceID:   resourceID,
	}
}

// String returns g in string format.
func (g GRID) String() string {
	return fmt.Sprintf("grid:%s:%s:%s:%s", g.Namespace, g.Owner, g.ResourceType, g.ResourceID)
}

// Parse attempts to parse a GRID from s.
func Parse(s string) (GRID, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 5 {
		return GRID{}, badFormatError(s, "there should be exactly 5 sections")
	} else if parts[0] != "grid" {
		return GRID{}, badFormatError(s, "GRIDs must start with 'grid:'")
	}

	return New(parts[1], parts[2], parts[3], parts[4]), nil
}

func badFormatError(s, reason string) error {
	return fmt.Errorf("'%s' is not in GRID format (%s); %s", s, Format, reason)
}
