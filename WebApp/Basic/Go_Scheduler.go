package main

import (
	"fmt"
	"runtime"
)

func main() {
	//G-goroutines - in user space threads scheduler
	//M-Machine worker threads - kernel threads space entity scheduler
	//P-Processor context - global-local runq time
	numCPU := runtime.NumCPU() // logical CPU in machine
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(2)
	numP1:=runtime.GOMAXPROCS(0)
	fmt.Println(numP1)
}