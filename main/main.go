package main

import (
	"fmt"

	"github.com/alwaysday1/twittercharacter"
)

func main() {
	postContent := "12345678"

	length := twittercharacter.CountTweet(postContent)

	fmt.Printf("推文长度根据自定义规则：%d\n", length)
}
