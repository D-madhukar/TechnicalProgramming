package main
import ("fmt"
		"time")
func hello(done chan int){
	fmt.Println("hello world")
	done <- 1
} 
func square(number int,schan chan int){
	time.Sleep(200000)
	schan<-number*number
}
func cube(number int,cchan chan int){
	cchan<-number*number*number
}
func main(){
	done:=make(chan int)
	go hello(done)
	<-done
	fmt.Println("Main function")

	schan:=make(chan int)
	cchan:=make(chan int)
	go square(5,schan)
	go cube(5,cchan)
	squares,cubes:=<-schan,<-cchan
	fmt.Println(squares,"   ",cubes)
	
}