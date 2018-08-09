# Global Resource ID

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/grid/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpatrick/grid)](https://goreportcard.com/report/github.com/zpatrick/grid)
[![Go Doc](https://godoc.org/github.com/zpatrick/grid?status.svg)](https://godoc.org/github.com/zpatrick/grid)

## Overview
GRID is a package that can be used to uniquely identify resources using Global Resource IDs (GRIDs). 
The GRID format is:
```
grid:namespace:object_type:object_id
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
        g, err := grid.Parse("grid:facebook:comment:53881")
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("GRID:       ", g.String())
        fmt.Println("Namespace:  ", g.Namespace)
        fmt.Println("Object Type:", g.ObjectType)
        fmt.Println("Object ID:  ", g.ObjectID)
}
```

Output:
```console
$ go run main.go
GRID:        grid:facebook:comment:53881
Namespace:   facebook
Object Type: comment
Object ID:   53881
```

## License
This work is published under the MIT license.

Please see the [LICENSE](/LICENSE) file for details.
