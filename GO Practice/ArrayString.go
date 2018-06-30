package main
import ("fmt"
		"strings")

func getAverage(arr []float32,size int) float32{
	var sum float32=0.0
	for i:=0;i<size;i++{
		sum=sum+arr[i]
	}
	return sum/float32(size)
}
// func main(){
// 	// var firstArray[10] float32
// 	var firstArray=[]float32{1.21,21.21,121.121,12.32,3.12}
// 	var stringArray=[]string{"madhu","madhukar"}
// 	fmt.Println(strings.Join(stringArray,","))
// 	for i:=0;i<len(firstArray);i++{
// 		fmt.Println(firstArray[i])
// 	}
// 	var twoD=[2][2]int{{2,0},{343,43}}//TOW DIMENSIONAL ARRAY
// 	for i:=0;i<2;i++{
// 		for j:=0;j<2;j++{
// 			fmt.Print(twoD[i][j],"  ")
// 		}
// 	}
// 	fmt.Println(getAverage(firstArray,5))
// 	var stringPointer string="this is pointer to a string"
// 	fmt.Println(stringPointer)

// 	var sliceArray []int
// 	sliceArray=make([]int,5,5)
// 	fmt.Println(len(sliceArray),"****",cap(sliceArray))

// 	var rangeArray=[]int{1,2,3,4,5,6}
// 	for i:=range rangeArray{
// 		fmt.Println(rangeArray[i],"   ")

// 	}

// 	var myMap map[string] string
// 	myMap=make(map[string] string)
// 	myMap["a"]="abc"
// 	myMap["b"]="scs"
// 	myMap["b"]="abcde"//UNIQUE KEYS
// 	for key:=range myMap{
// 		fmt.Println(key,"=",myMap[key])
// 	}
// }