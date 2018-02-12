package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get("http://example.com/").
		RedirectPolicy(redirectPolicyFunc).
		Set("If-None-Match", `W/"wyzzy"`).
		End()
}
