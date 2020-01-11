# i18n for Go

A simple i18n library that uses yaml to define localization files.

## Features

- Hierarchical model (Based on the Yaml)
- Efficient access through grouping
- Cache created instances
- Cache the last used instance

## Usage

**Please refer to the `example` package**

```yaml
user:
  password:
    error:
      too-short: "密码太短: %s"
      too-simple: 密码不符合规则
```

```go
package main

import (
  "github.com/eightpigs/i18n"
  "fmt"
)

func main() {
  // Create instances based on default locale
  locale, e := i18n.New()
  if e != nil {
  	panic(e)
  }
  
  // get message
  msg := locale.Get("user.password.error.too-simple")
  fmt.Println(msg)
  
  // get message group
  group := locale.Group("user.password.error")
  
  // get messages within a group
  msg = group.Get("too-short", "12345")
  fmt.Println(msg)
}
```

Specify localization language

```go
locale, e := i18n.NewLocale("zh-CN", "xxx/locales/zh-CN.yaml")
```

Sets the global default localization language

```go
i18n.DefaultLocale = "zh-CN"
// Create instances based on default locale
i18n.New()
```

Use the last localize instance.

```go
// Use the last localize instance.
msg = i18n.Get("user.password.error.too-simple")
fmt.Println(msg)

group = i18n.Group("user.password.error")
fmt.Println(group.Get("too-short", "112233"))
```

## Benchmark

`go test -bench=. -benchmem -benchtime=10s -count=2 -run=none` on

- CPU: i7-9700K (8) @ 4.900GHz

```
goos: linux
goarch: amd64
pkg: github.com/eightpigs/i18n
Benchmark_Normal-8   	578327178	        20.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_Normal-8   	585156315	        20.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_Parse-8    	84822812	       143 ns/op	      40 B/op	       2 allocs/op
Benchmark_Parse-8    	79351177	       147 ns/op	      40 B/op	       2 allocs/op
PASS
ok  	github.com/eightpigs/i18n	52.125s
```

## LICENSE

i18n is available under the MIT license. See the [LICENSE](https://github.com/eightpigs/i18n/blob/master/LICENSE) file for more info.
