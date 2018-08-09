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

	fmt.Println("GRID:       ", g.String())
	fmt.Println("Namespace:  ", g.Namespace)
	fmt.Println("Owner:      ", g.Owner)
	fmt.Println("Object Type:", g.ObjectType)
	fmt.Println("Object ID:  ", g.ObjectID)
	// Output:
	// GRID:        grid:facebook:user123:comment:88532
	// Namespace:   facebook
	// Owner:       user123
	// Object Type: comment
	// Object ID:   88532
}

func TestParse(t *testing.T) {
	cases := map[string]GRID{
		"grid:namespace:owner:object_type:object_id": New("namespace", "owner", "object_type", "object_id"),
		"grid::owner:object_type:object_id":          New("", "owner", "object_type", "object_id"),
		"grid:::object_type:object_id":               New("", "", "object_type", "object_id"),
		"grid::::object_id":                          New("", "", "", "object_id"),
		"grid::::":                                   New("", "", "", ""),
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
		"missing 4 sections":        "object_id",
		"missing 3 sections":        "object_type:object_id",
		"missing 2 sections":        "owner:object_type:object_id",
		"missing 2 section":         "namespace:owner:object_type:object_id",
		"doesn't start with 'grid'": "rgid:namespace:object_type:object_id",
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := Parse(input); err == nil {
				t.Fatalf("Error was nil!")
			}
		})
	}
}
