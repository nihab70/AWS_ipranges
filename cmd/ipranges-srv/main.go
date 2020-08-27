package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nihab70/cloudorama/cloudkit/aws"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("http service stating...")
	http.HandleFunc("/", returnIPRange)
	http.HandleFunc("/status", returnStatus)
	http.ListenAndServe(":8080", nil)
	log.Info("http service shuting down...")
}

func returnIPRange(w http.ResponseWriter, r *http.Request) {

	data := aws.GetAWSIPRange()

	e, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	args := r.URL.Path[1:]

	if args == "" {
		fmt.Fprintf(w, "Result: \n %s", string(e))
	}
}

func returnStatus(w http.ResponseWriter, r *http.Request) {

	data := aws.GetAWSIPRange()

	status := string(data.CreateDate) + "\n" + string(data.SyncToken)

	fmt.Fprintf(w, "Status:\n%s", status)
}
