package main
import "fmt"
func main(){
	var a int=10
	var b int=0;
	for a>b{
		fmt.Printf("a=%d is still >b=%d\n",a,b)
		fmt.Println("a=",a,"is still > b=",b)
		b++
	}
	for i:=0;i<10;i++{
		fmt.Println("think think")
	}
	switch a{
		case 9:fmt.Println("hey it is 9")
		case 10:fmt.Println("No it is 10 only")
	default:fmt.Println("hahaha neither 10 nor 9")
	}
}