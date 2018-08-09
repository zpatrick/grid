# Global Resource ID

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/grid/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpatrick/grid)](https://goreportcard.com/report/github.com/zpatrick/grid)
[![Go Doc](https://godoc.org/github.com/zpatrick/grid?status.svg)](https://godoc.org/github.com/zpatrick/grid)

## Overview
GRID is a package that can be used to uniquely identify resources using Global Resource IDs (GRIDs). 
The GRID format is:
```
grid:namespace:owner:resource_type:resource_id
```

## Download
To download this package, run:
```
go get github.com/zpatrick/grid
```

## Usage
This code is taken from the [example](/example) located in this repo.

```go
package main

import (
        "fmt"
        "log"

        "github.com/zpatrick/grid"
)

func main() {
	g, err := grid.Parse("grid:facebook:user123:comment:88532")
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("GRID:         ", g.String())
        fmt.Println("Namespace:    ", g.Namespace)
        fmt.Println("Owner:        ", g.Owner)
        fmt.Println("Resource Type:", g.ResourceType)
        fmt.Println("Resource ID:  ", g.ResourceID)
}
```

Output:
```console
$ go run main.go
GRID:          grid:facebook:comment:88532
Namespace:     facebook
Owner:         user123
Resource Type: comment
Resource ID:   88532
```

## License
This work is published under the MIT license.

Please see the [LICENSE](/LICENSE) file for details.
