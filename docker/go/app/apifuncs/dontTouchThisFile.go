package apifuncs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"set1.ie.aitech.ac.jp/HackU_vol_1/dbctl"
)

// DBTest is func that tests database funcs.
func DBTest(w http.ResponseWriter, r *http.Request) {
	// CORS settings.
	w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)              // Allow remote access.
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Allowed methods.

	if r.Method == http.MethodGet {
		// CallTasks returns []dbctl.Task.
		tasks, err := dbctl.CallTasks()
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		// Array to JSON.
		b, err := json.Marshal(tasks)
		if err != nil {
			log.Fatal(err)
		}
		// []byte to string.
		taskString := string(b)

		w.WriteHeader(200)
		fmt.Fprintln(w, taskString)
		fmt.Println(taskString)
	}
}
