# WRR

Weighted Round Robin Balancer

---

## Getting Started

```
go get -u github.com/yoonper/wrr
```

```
package main

import (
	"fmt"
	"github.com/yoonper/wrr"
)

func main() {
	b := wrr.Init()
	b.Add("item1", 2)
	b.Add("item2", 5)
	b.Add("item3", 3)
	count := make(map[string]int)
	for i := 0; i < 100; i++ {
		item := b.Next()
		count[item]++

		// you can increase or decrease item weight
		// b.IncWeight("item3")
		// b.DecWeight("item3")
	}
	fmt.Println(count)
}

```

---

## Reference

[ngx_http_upstream_round_robin.c](https://github.com/phusion/nginx/blob/master/src/http/ngx_http_upstream_round_robin.c)
