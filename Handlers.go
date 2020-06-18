package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

// Handles "/"
func HandleMain(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Reached HandleMain")
}

// Handles "/index"
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index!")
}

// Handles "/userId/{id}"
func HandleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	fmt.Fprintln(w, "User ID:", userId)
}

// Handles "/data-dump"
func HandleDump(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	_, err = os.Open("filename.txt")
	if err != nil {
		logger := &logrus.Logger{
			Out:   os.Stdout,
			Level: logrus.DebugLevel,
			Formatter: &easy.Formatter{
				TimestampFormat: "2006-01-02 15:04:05",
				LogFormat:       "[ERROR]: %time% - %msg% \n",
			},
		}
		logger.Println("Error Encountered:", stackTrack, err, http.StatusConflict)

		data, _ := ioutil.ReadAll(response.Body)
		fmt.Fprintln(w, string(data))
	}
}

func Swagger(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./swaggerui/index.html"))
}
