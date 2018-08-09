package main

import (
	"fmt"
	"log"

	"github.com/zpatrick/grid"
)

func main() {
	g, err := grid.Parse("grid:facebook:comment:53881")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GRID:       ", g.String())
	fmt.Println("Namespace:  ", g.Namespace)
	fmt.Println("Object Type:", g.ObjectType)
	fmt.Println("Object ID:  ", g.ObjectID)
}
