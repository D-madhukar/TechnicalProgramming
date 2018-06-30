package main
import (
    "fmt"
)
func changeArray(s []int){
    s[0]=10
    return
}
func main(){
  var arr []int
  arr=[]int{1,2,3,4}
  slicearr:=arr[:]
  // changeArray(slicearr)
  changeArray(arr)
  fmt.Println(arr)
  fmt.Println(slicearr)
}
