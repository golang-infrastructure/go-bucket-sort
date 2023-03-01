# 桶排序（Bucket Sort）的Golang实现

# 一、这个是什么？优势是啥？

这是桶排序的一个Golang实现，优势如下：

- 支持对负数进行桶排序
- 将所需的bucket压缩到尽可能小，避免空间浪费，空间利用率高
- 支持bucket数量限制，避免意外情况OOM造成应用Crash

# 二、安装

```bash
go get -u github.com/golang-infrastructure/go-bucket-sort
```

# 三、实例代码

```go
package main

import (
	"fmt"
	bucket_sort "github.com/golang-infrastructure/go-bucket-sort"
	"math/rand"
)

func main() {
	slice := make([]int, 10)
	for index := range slice {
		slice[index] = rand.Intn(100)
		if rand.Int()%2 == 0 {
			slice[index] *= -1
		}
		//slice[index] *= -1
	}
	err := bucket_sort.Sort(slice, 1000)
	fmt.Println(err)
	fmt.Println(slice)
	// Output:
	// <nil>
	// [-81 -62 -25 11 28 37 47 56 81 94]
}
```

# 四、TODO

- 完善文档，加上对任意类型数组进行桶排序的扩展的API示例和文档 
- 桶排序应该是需要稳定的 









