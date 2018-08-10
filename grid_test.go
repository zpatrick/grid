package grid

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleGRID() {
	g := New("facebook", "user123", "comment", "88532")
	fmt.Println(g.String())
	// Output:
	// grid:facebook:user123:comment:88532
}

func ExampleParse() {
	g, err := Parse("grid:facebook:user123:comment:88532")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GRID:         ", g.String())
	fmt.Println("Namespace:    ", g.Namespace)
	fmt.Println("Owner:        ", g.Owner)
	fmt.Println("Resource Type:", g.ResourceType)
	fmt.Println("Resource ID:  ", g.ResourceID)
	// Output:
	// GRID:          grid:facebook:user123:comment:88532
	// Namespace:     facebook
	// Owner:         user123
	// Resource Type: comment
	// Resource ID:   88532
}

func TestParse(t *testing.T) {
	cases := map[string]GRID{
		"grid:namespace:owner:resource_type:resource_id": New("namespace", "owner", "resource_type", "resource_id"),
		"grid::owner:resource_type:resource_id":          New("", "owner", "resource_type", "resource_id"),
		"grid:::resource_type:resource_id":               New("", "", "resource_type", "resource_id"),
		"grid::::resource_id":                            New("", "", "", "resource_id"),
		"grid::::":                                       New("", "", "", ""),
	}

	for s, expected := range cases {
		t.Run(s, func(t *testing.T) {
			result, err := Parse(s)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expected, result)
		})
	}
}

func TestParseError(t *testing.T) {
	cases := map[string]string{
		"missing 5 sections":        "",
		"missing 4 sections":        "resource_id",
		"missing 3 sections":        "resource_type:resource_id",
		"missing 2 sections":        "owner:resource_type:resource_id",
		"missing 1 section":         "namespace:owner:resource_type:resource_id",
		"doesn't start with 'grid'": "rgid:namespace:resource_type:resource_id",
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := Parse(input); err == nil {
				t.Fatalf("Error was nil!")
			}
		})
	}
}
