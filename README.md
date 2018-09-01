# 📶 emoji-ordering
> Utilities for ordering emojis according to the official Unicode emoji sort order

[![Go Report Card](https://goreportcard.com/badge/github.com/bcongdon/emoji-ordering)](https://goreportcard.com/report/github.com/bcongdon/emoji-ordering)
[![GoDoc](https://godoc.org/github.com/bcongdon/emoji-ordering?status.svg)](https://godoc.org/github.com/bcongdon/emoji-ordering)

## Installation

```
go get github.com/bcongdon/emoji-ordering
```

## Usage

```go
package main

import (
    "github.com/bcongdon/emoji-ordering"
    "fmt"
    "sort"
)

func main() {
    emojis := []string{"😎", "💯", "🤔", "🤷‍♂️", "🤷"}
    sort.Sort(ordering.EmojiSlice(emojis))
    fmt.Println(emojis)
    // [🤔, 😎, 💯, 🤷, 🤷‍♂️]

    fmt.Println(ordering.IsEmoji("🤷‍♂️"))
    // true

    fmt.Println(ordering.IsEmoji("foo"))
    // false

    fmt.Println(ordering.IsEmoji("😀 ")) // with extra whitespace
    // false

    // Constructing an Emoji Slice with Validation
    emojis = []string{"😎", "💯", "foo"}
    _, err := ordering.NewEmojiSlice(emojis)
    fmt.Println(err)
    // 'foo' is not an emoji
}
```