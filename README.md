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
	"eightpigs.io/i18n"
	"fmt"
)

func main() {
  // Load the instance from ./locales/zh-CN.yaml
  locale, e := i18n.NewLocale("zh-CN", "")
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
