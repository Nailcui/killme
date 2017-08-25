# Kill me
---
Kill me allow you to kill the program you just run more easily. 
Kill me 可以更方便的关闭在后台运行的程序。

Installation
---
```
go get github.com/nailcui/killme
```

Overview
---
像下面例子：main.go中一样导入killme:
```
package main

import (
	"fmt"
	_ "github.com/nailcui/killme"
	"time"
)

func main() {
	fmt.Println("sleep...")
	time.Sleep(10 * time.Minute)
	fmt.Println("exit...")
}
```
然后在一个控制台中运行：
```
go run main.go
```
此时会在运行目录下生成`killme.sh`文件，运行此文件命令将会调用`kill -9 pid`将刚才的程序kill掉。
接着打开另一个控制台运行这个`killme.sh`
```
./killme.sh
```
此时可以看到刚才的main程序被关掉了

Examples
---
[Basic](https://github.com/Nailcui/killme/blob/master/examples/main.go)
