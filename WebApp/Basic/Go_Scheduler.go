package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	//"time"
)
var wg sync.WaitGroup

func g1(){fmt.Println("go 1")
wg.Done()}
func g2(){
	fmt.Println("go 2")
    os.OpenFile("hello", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

//logic 1
//logic 2
wg.Done()
}
func g3(){fmt.Println("go 3")
wg.Done()}
func g4(){fmt.Println("go 4")
wg.Done() }
func main() {
	//G-goroutines - in user space threads scheduler
	//M-Machine worker - OS threads - kernel threads space entity scheduler
	//P-Processor context - global-local runq time
	numCPU := runtime.NumCPU() // logical CPU in machine
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(2) // set 2 kernel CPU in machine for P-local runq-concurrency-context switch
	numP1:=runtime.GOMAXPROCS(0)
	fmt.Println(numP1)
	//concurrency-context switch
	//parallelism
	wg.Add(4)
	fmt.Println("...begin")
	go g1() //Gmain->P1
	go g2() //->P1-->M1,P context switch
	go g3() //Gmain->P2
	go g4() //->P2-->M2->g1,g2 remove when completed, context switch
	//os. open file ->sys call->locked->create new M3, split M1 run completed (wait,reuse, sys call locked-link - thread spinning to P1)
	//->P context switch->M2
	//dynamic optimized memory - resource CPU
	// all P-M - teminated when all Goroutines completed

	wg.Wait()

	//time.Sleep(time.Second)
	fmt.Println("...end")

}