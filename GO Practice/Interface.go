package main
import "fmt"
	
func main(){
	myInt:=IntCounter(0)
	var myIncrementer Incrementer=&myInt
	fmt.Println(myIncrementer.Increment())
}
type Incrementer interface{
	 Increment() int
}
type IntCounter int
func (cn *IntCounter) Increment() int{
	*cn++
	return int(*cn)
}