package main

import(
	"fmt"
	"net/http"	
	"set1.ie.aitech.ac.jp/HackU_vol_1/apifuncs"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		apifuncs.Test(w,r)
		
	  fmt.Fprintln(w, "Hello World")	  	  
	})

	http.ListenAndServe(":8080", nil)	
  }

