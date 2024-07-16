# Twitter Character Counter

A Go package to accurately count the length of a tweet according to Twitter's character counting rules.

## Features

- Accurately counts tweet length based on Twitter's rules
- Handles URLs and special characters correctly
- Easy to integrate into existing Go projects

## Installation
```sh
go get github.com/alwaysday1/twitter-character-counter
```
## Usage

```go
package main

import (
    "fmt"
    "github.com/alwaysday1/twitter-character-counter"
)

func main() {
    postContent := "12345678"
    length := twittercharacter.CountTweet(postContent)
    fmt.Printf("Tweet length: %d\n", length)
}
```
## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.