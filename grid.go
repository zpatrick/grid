package grid

import (
	"fmt"
	"strings"
)

// Format of all GRIDs
const Format = "grid:namespace:object_type:object_id"

// GRID (Global Resource ID) objects are used to uniquely identify resources.
type GRID struct {
	Namespace  string
	ObjectType string
	ObjectID   string
}

// New returns a new GRID object.
func New(namespace, objectType, objectID string) GRID {
	return GRID{
		Namespace:  namespace,
		ObjectType: objectType,
		ObjectID:   objectID,
	}
}

// String returns g in string format.
func (g GRID) String() string {
	return fmt.Sprintf("grid:%s:%s:%s", g.Namespace, g.ObjectType, g.ObjectID)
}

// Parse attempts to parse a GRID from s.
func Parse(s string) (GRID, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 4 {
		return GRID{}, badFormatError(s, "there should be exactly 4 sections")
	} else if parts[0] != "grid" {
		return GRID{}, badFormatError(s, "GRIDs must start with 'grid:'")
	}

	return New(parts[1], parts[2], parts[3]), nil
}

func badFormatError(s, reason string) error {
	return fmt.Errorf("'%s' is not in GRID format (%s); %s", s, Format, reason)
}
