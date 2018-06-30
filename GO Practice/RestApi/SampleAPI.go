package main
import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)
type Student struct{
	ID 		string		`json:"id"`	
	Name	string		`json:"name,omitempty"`
	Phoneno string		`json:"phoneno,omitempty"`
}
var students []Student
func sampleHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Hello world")
	fmt.Fprintf(w,"<form action='http://localhost:3132/poststudents'><input type='submit'></form>")
}
func getStudentsDetails(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(students)
	//fmt.Fprintf(w,"<form action='http://localhost:3132/poststudents'><input type='submit'></form>")
}
func getStudent(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	var id string =params["id"]
	// fmt.Println(id,"***")
	for _,item:=range students{
		fmt.Printf("%T**%T",id,item.ID)
		if id==item.ID{
			json.NewEncoder(w).Encode(item)
			fmt.Println(item,"***")
			break
		}
	}
}
var students1 []Student
func postStudent(w http.ResponseWriter,r *http.Request){
	json.NewDecoder(r.Body).Decode(&students1)
	json.NewEncoder(w).Encode(students1)
	fmt.Println(students1)
}
func main(){
	router:=mux.NewRouter()
	students=append(students,Student{ID:"100",Name:"madhukar",Phoneno:"12345678"})
	students=append(students,Student{ID:"101",Name:"sanju",Phoneno:"12345678"})
	router.HandleFunc("/",sampleHandler).Methods("GET")
	router.HandleFunc("/students",getStudentsDetails).Methods("GET")
	router.HandleFunc("/student/{id}",getStudent).Methods("GET")
	router.HandleFunc("/poststudents",postStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":3132",router))
}