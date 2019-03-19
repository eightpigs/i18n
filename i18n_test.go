package i18n

import (
	"testing"
)

func Test_Get(t *testing.T) {
	locale, e := NewLocale("zh-CN", "example/locales/zh-CN.yaml")
	if e != nil {
		panic(e)
	}
	g := locale.Group("user.password.error")
	msg := g.Get("too-short", "1234a")
	expected := "密码太短: 1234a"
	if msg != expected {
		t.Errorf("i18n format error: %s != %s", msg, expected)
	} else {
		t.Logf("%#v\n", msg)
	}
}

func Benchmark_Normal(b *testing.B) {
	locale, e := NewLocale("zh-CN", "example/locales/zh-CN.yaml")
	if e != nil {
		panic(e)
	}

	g := locale.Group("user.password.error")
	for i := 0; i < b.N; i++ {
		_ = g.Get("too-simple")
	}
}

func Benchmark_Parse(b *testing.B) {
	locale, e := NewLocale("zh-CN", "example/locales/zh-CN.yaml")
	if e != nil {
		panic(e)
	}

	g := locale.Group("user.password.error")
	for i := 0; i < b.N; i++ {
		_ = g.Get("too-short", b.N)
	}
}
