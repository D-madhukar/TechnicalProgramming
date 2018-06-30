package main
import (
	"fmt"
	"time"
)
func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)
	go func(){
		for{
			ch1<-"hello--500ms"
			time.Sleep(time.Millisecond*500)
		}
	}()
	go func(){
		for {
			ch2<-"world--2sec"
			time.Sleep(time.Second*2)
		}
		close(ch2)
	}()
	for{
		select{
			case m1:=<-ch1 : fmt.Println(m1)
			case m2,ok:=<-ch2 :{ 
				if ok==true{
						fmt.Println(m2)
				}
			}
		}
	}
}