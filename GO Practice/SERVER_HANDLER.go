package main
import ("fmt"
		"net/http")
func input_handle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"<h1>hello World</h1>")
} 		
func main(){
	http.HandleFunc("/",input_handle)
	http.ListenAndServe(":3131",nil)
}