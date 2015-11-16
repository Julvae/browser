// Regexp project main.go
package main

import (
	"Regexp/tools"
)

//目前开多个协程还不行，运行的时候把输入锁掉
func main() {
	tools.Run()
}
