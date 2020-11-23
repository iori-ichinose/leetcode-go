package design

import (
	"fmt"
	"math/rand"
	"strings"
)

const hashnums string = "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type Codec struct {
	urls map[string]string
}

func TinyURLConstructor() Codec {
	return Codec{urls: map[string]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	s := make([]string, 6)
	for i := 0; i < 6; i++ {
		s[i] = string(hashnums[rand.Intn(62)])
	}
	shortUrl := "tiny.url/" + strings.Join(s, "")
	_, y := this.urls[shortUrl]
	if !y {
		this.urls[shortUrl] = longUrl
		return shortUrl
	} else {
		return this.encode(longUrl)
	}
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.urls[shortUrl]
}

func TinyURLTb() {
	fmt.Println("Testbench TinyURL: ")
	c := TinyURLConstructor()
	long := []string{"www.baidu.com/nmsl", "www.baidu.com/wdnmd", "www.baidu.com/fnmdp"}
	short := []string{}
	after := []string{}
	for _, n := range long {
		s := c.encode(n)
		short = append(short, s)
		after = append(after, c.decode(s))
	}
	for n := range long {
		println(long[n], short[n], after[n])
	}
}
