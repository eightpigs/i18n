package main

import (
	"fmt"
	"github.com/eightpigs/i18n"
)

func main() {
	locale, e := i18n.NewLocale("zh-CN", "example/locales/zh-CN.yaml")
	//i18n.DefaultLocale = "zh-CN"
	//locale, e := i18n.New()
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

	// Use the last localized configuration.
	msg = i18n.Get("user.password.error.too-simple")
	fmt.Println(msg)

	group = i18n.Group("user.password.error")
	fmt.Println(group.Get("too-short", "112233"))
}
