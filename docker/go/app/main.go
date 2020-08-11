package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type InputJsonSchema struct{

	Name string `json:"name"`
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "Get Method")
		
		  } else if r.Method== http.MethodPost{
			fmt.Fprintln(w, "Post Method")

			fmt.Fprintln(w,r.FormValue("in"))
			fmt.Fprintln(w,r.FormValue("name"))
			fmt.Fprintln(w,r.FormValue("required"))
			fmt.Fprintln(w,r.FormValue("schema"))
			
	
		  }else if r.Method==http.MethodPut{
			fmt.Fprintln(w,"Put Method")  			
		  }else{
			fmt.Fprintln(w,"delete Method")			
		  }
						  
	  fmt.Fprintln(w, "Hello World")	  	  
	})

	http.ListenAndServe(":8080", nil)	
  }

