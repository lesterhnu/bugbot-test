package main

import (
	"fmt"

	"bugbot/cases/goroutine_leak"
	"bugbot/cases/mysql_unclosed"
	"bugbot/cases/nil_pointer"
	"bugbot/cases/slice_oob"
)

func main() {
	fmt.Println("buggy code demo started")

	goroutineleak.Run()
	sliceoob.Run()
	nilpointer.Run()
	mysqlunclosed.Run()

	fmt.Println("buggy code demo finished")
}
