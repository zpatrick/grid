package grid

import (
	"fmt"
	"strings"
)

// GRID (Global Resource ID) uniquely identify resources.
// the GRID format is
//    grid:namespace:object_type:object_id
type GRID struct {
	Namespace  string
	ObjectType string
	ObjectID   string
}

// Returns a new GRID object
func New(namespace, objectType, objectID string) GRID {
	return GRID{
		Namespace:  namespace,
		ObjectType: objectType,
		ObjectID:   objectID,
	}
}

// Returns a string in GRID in format
func (g GRID) String() string {
	return fmt.Sprintf("grid:%s:%s:%s", g.Namespace, g.ObjectType, g.ObjectID)
}

// Attempts to parse a GRID from s
func Parse(s string) (GRID, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 4 || parts[0] != "grid" {
		return GRID{}, fmt.Errorf("'%s' is not in GRID format, there should be exactly 4 sections", s)
	}

	return New(parts[1], parts[2], parts[3]), nil
}
