package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Infof("GET request")

	if dish, ok := r.URL.Query()["m"]; ok && len(dish) > 0 {
		if serve, ok := r.URL.Query()["s"]; ok && len(serve) > 0 {
			if intVar, err := strconv.Atoi(serve[0]); err == nil { // convert string to int
				// call function <<ingredients>> in python module <<pt.py>>
				// passing arguments
				cmd := exec.Command("python3", "-c", fmt.Sprintf("import pt; print(pt.ingredients('%v', %v))", dish[0], intVar))
				fmt.Println(cmd.Args)
				// retrieve pthon function return values
				if out, err := cmd.CombinedOutput(); err == nil {
					// replace ' with "
					// remove \n from the end of string
					ss := strings.ReplaceAll(strings.TrimSpace(string(out)), `'`, `"`)
					fmt.Println(ss)
					var resp map[string]string
					if err := json.Unmarshal([]byte(ss), &resp); err == nil {
						if err := json.NewEncoder(w).Encode(resp); err != nil {
							http.Error(w, err.Error(), 500)
						}
					} else {
						fmt.Println(err)
					}
					return
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		}
	}
	if err := json.NewEncoder(w).Encode(map[string]string{"from": `T-T`, "status": "ALL GOOD"}); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
