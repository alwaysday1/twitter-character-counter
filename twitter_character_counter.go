package common

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	TWITTER_URL_SIZE            = 23
	TWITTER_SPECIAL_CHAR_SIZE   = 2
	TWITTER_STANDARD_CHAR_LIMIT = 0x2037
)

var (
	URL_MATCH = regexp.MustCompile(`https?://[^\s]+`)
)

// isSpecialChar 判断字符是否为表情符号或其他非标准的 Unicode 字符
func isSpecialChar(r rune) bool {
	if r >= 0xFE00 && r <= 0xFE0F {
		return false // 忽略变体选择符
	}
	return unicode.Is(unicode.S, r) || unicode.Is(unicode.M, r) || unicode.Is(unicode.P, r) || unicode.Is(unicode.Sk, r)
}

// findSpecialChars 查找文本中的所有表情符号和其他非标准的 Unicode 字符
func findSpecialChars(text string) []string {
	var specialChars []string
	for _, char := range text {
		if isSpecialChar(char) && int(char) > TWITTER_STANDARD_CHAR_LIMIT {
			specialChars = append(specialChars, string(char))
		}
	}
	return specialChars
}

// findUrls 查找文本中的所有 URL，包括没有 "http://" 前缀的 URL
func findUrls(text string) []string {
	return URL_MATCH.FindAllString(text, -1)
}

// CountTweet 根据推特的规则计算推文的总长度
func CountTweet(tweetText string) int {
	// 查找并移除文本中的 URL
	urls := findUrls(tweetText)
	for _, url := range urls {
		tweetText = strings.Replace(tweetText, url, "", 1)
	}

	// 查找并移除文本中的特殊字符
	specialChars := findSpecialChars(tweetText)
	for _, char := range specialChars {
		tweetText = strings.Replace(tweetText, char, "", 1)
	}

	// 计算推文的总长度
	urlsLength := len(urls) * TWITTER_URL_SIZE
	specialCharsLength := len(specialChars) * TWITTER_SPECIAL_CHAR_SIZE
	tweetLength := utf8.RuneCountInString(tweetText) + urlsLength + specialCharsLength

	return tweetLength
}
