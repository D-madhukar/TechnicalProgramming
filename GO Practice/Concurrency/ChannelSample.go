package main
import "fmt"
func producer(ch chan int){
	for i:=0;i<10;i++{
		ch<-i
	}
	close(ch)
}
// func main(){
// 	ch:=make(chan int)
// 	go producer(ch)
// 	for{
// 		v,ok:=<-ch
// 		if ok==false{
// 			fmt.Println("channel is closed")
// 			break
// 		}
// 		fmt.Println("ok==",ok,"v=",v)
// 	}
// 	ch=make(chan int)
// 	go producer(ch)
// 	for v:=range ch{
// 		fmt.Println(v)
// 	}

// }