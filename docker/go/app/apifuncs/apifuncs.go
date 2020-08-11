package apifuncs

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
)


type InputJsonSchema struct{

	In string `json:"in"`
	Name string `json:"name"`
	Description string `json:"desciption"`
	Required bool `json:"required"`
	Responses string `schema:"responses"`
}

func Test(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		Jsontype:=InputJsonSchema{"in","Name","desciption",true,"responses"}

		fmt.Fprintln(w,Jsontype)
		
		fmt.Fprintln(w, "Get Method")
		
			
	  } else if r.Method== http.MethodPost{
		fmt.Fprintln(w, "Post Method")


		b, err:=ioutil.ReadAll(r.Body)
		if err !=nil{
			fmt.Println("io error")
			return 
		}

		jsonBytes := ([]byte)(b)

		data:=new(InputJsonSchema)
		if err:=json.Unmarshal(jsonBytes,data);err!=nil{
			fmt.Println("JSON Unmarshal error:",err)
			return
		}

		fmt.Fprintf(w,"jsonを受け取りました")
/* 
		fmt.Fprintf(w,data.In)
		fmt.Fprintf(w,data.Name)
		fmt.Fprintf(w,data.Description)
		fmt.Fprintf(w,data.Required)
		fmt.Fprintf(w,data.Responses)	 */			
										
	  }else if r.Method==http.MethodPut{
		fmt.Fprintln(w,"Put Method")  			
	  }else{
		fmt.Fprintln(w,"delete Method")			
	  }
}
