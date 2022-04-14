package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	log.Infof("GET request")

	cmd := exec.Command("python", "-c", "import pt; print(pt.ingredients('foo', 'bar'))")
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(map[string]string{"from": `T-T`, "status": "ALL GOOD"}); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
