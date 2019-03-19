# i18n for Go

A simple i18n library that uses yaml to define localization files.

## Features

- Hierarchical model (Based on the Yaml)
- Efficient access through grouping
- Cache created instances

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

## Benchmark

```
goos: darwin
goarch: amd64
pkg: eightpigs.io/i18n
Benchmark_Normal-8   	30000000	        44.2 ns/op
Benchmark_Parse-8    	 5000000	       247 ns/op
PASS
```

## LICENSE

i18n is available under the MIT license. See the [LICENSE](https://github.com/eightpigs/i18n/blob/master/LICENSE) file for more info.
