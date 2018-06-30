package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
)
type Student struct{
	ID 		string		`json:"id"`	
	Name	string		`json:"name,omitempty"`
	Phoneno string		`json:"phoneno,omitempty"`
}
func main(){
	var students []Student
	res,err:=http.Get("http://localhost:3132/students")
	if err!=nil{
		fmt.Println("error in getting response")
	}
	body,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("error in JSON")
	}
	err=json.Unmarshal(body,&students)
	if err!=nil{
		fmt.Println("*ERROR in unmarshaling")
	}
	fmt.Println(students)
	jsonvalue,_:=json.Marshal(students)
	response,err1:=http.Post("http://localhost:3132/poststudents","application/json",bytes.NewBuffer(jsonvalue))
	if err1!=nil{
		fmt.Println("error",err1.Error())
	}else{
		data,_:=ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	
	}
	
}