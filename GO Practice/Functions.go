//Basic syntax of functions
//function parameters pass by value and pass by reference
//functions as values
//functions use ... to take multiple values of same type
//Annonymus functions
package main
import "fmt"
func swap(x,y string)(string,string){
	return y,x
}
func displayThisFunction(fun func(string) string){
	fmt.Println(fun("hell"))
}
func sum(values ...int){
	result:=0
	for _,v:=range values{//if you dont want to use variable(first before v) use underscore
		result+=v
	}
	fmt.Println(result)
}
func annonymousFunction(){
	var divide func(float32,float32) (float32,error)
	divide=func(x,y float32)(float32,error){
		if y==0{
			return 0.0,fmt.Errorf("Cannot divide by zero")
		}
		return x/y,nil
	}
	d,err:=divide(16.34,32.32)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
type Person struct{//Functions as Methods
	name string
	phoneno string
}
func (p Person) displayData(){
	fmt.Println(p.name,"***",p.phoneno)
}
func (p *Person) changeName(newName string){
	p.name=newName;
}
// func main(){
// 	// var c int=0 if any variable is  declared ...the program must contains
// 	//atleast one usage..It gives you run time error

// 	a,b:=swap("madhu","madhukar")
// 	fmt.Println(a,"   ",b)

// 	myfun1:=func(x string) string{
// 		return x+"world"
// 	}
// 	fmt.Println(myfun1("hello"))
// 	displayThisFunction(myfun1)
// 	displayThisFunction(func(x string) string{
// 		return x
// 	})
// 	sum(1,2,3,4,5,5,6)
// 	annonymousFunction()

// 	p1:=Person{ 
// 		name:"madhukar",
// 		phoneno:"8374676230",
// 	}
// 	p1.displayData()
// 	p1.changeName("madhu")
// 	fmt.Println(p1.name)
// }