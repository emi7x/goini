# GoIni [![Go Report Card](https://goreportcard.com/badge/github.com/mt6x/goini)](https://goreportcard.com/report/github.com/mt6x/goini)

GoIni is a minimalistic ini parser written in golang.

## Syntax

### Creating a new ini parser

When creating a new ini parser, the function returns two values; the ini parser, and an error value. To define a new ini parser;

```go
package main

import "github.com/mt6x/goini"

func main() {
    parser, err := goini.New(".", "path", "to", "file")
    if err != nil {
        // error handle here
    }
}
```

Once defined, it will parse the file without any further user input.

### Fetching an element from the file

If you want to fetch an element from the parsed file that is not in a section;

```go
package main

import "github.com/mt6x/goini"

func main() {
    parser, err := goini.New(".", "path", "to", "file")
    if err != nil {
        // error handle here
    }

    item := parser.GetValueFromKey("name-of-key")
    if item == nil {
        // error handle here
    }

    // Display the item as a string format instead of a string pointer.
    println(*item)
}
```

And if you want to fetch an item from a section;

```go
package main

import "github.com/mt6x/goini"

func main() {
    parser, err := goini.New(".", "path", "to", "file")
    if err != nil {
        // error handle here
    }

    item := parser.GetValueFromSectionKey("name-of-section", "name-of-key")
    if item == nil {
        // error handle here
    }

    // Display the item as a string format instead of a string pointer.
    println(*item)
}
```

### Reloading the ini parser

If you want to reload the ini parser;

```go
package main

import "github.com/mt6x/goini"

func main() {
    parser, err := goini.New(".", "path", "to", "file")
    if err != nil {
        // error handle here
    }

    if err := parser.ReloadKeys(); err != nil {
        // error handle here
    }

    // Display the item as a string format instead of a string pointer.
    println(*item)
}
```

## Features

GoIni supports both `#` and `;` comments, as well as ini sections. It will ignore empty lines as well as comments.
