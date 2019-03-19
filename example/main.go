package main

import (
	"eightpigs.io/i18n"
	"fmt"
)

func main() {
	locale, e := i18n.NewLocale("zh-CN", "")
	if e != nil {
		panic(e)
	}

	msg := locale.Get("user.password.error.too-simple")
	fmt.Println(msg)

	// get message group
	group := locale.Group("user.password.error")

	// get messages within a group
	msg = group.Get("too-short", "12345")
	fmt.Println(msg)
}