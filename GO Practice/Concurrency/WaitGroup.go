package main
import (
	"fmt"
	"sync"
)
func goRoutineCall(str string){
	fmt.Println(str)
}

// func main(){
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func(){
// 		goRoutineCall("hello")
// 		wg.Done()
// 	}()
// 	wg.Wait()
// 	fmt.Println("main exited")
// }