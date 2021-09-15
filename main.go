package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		currentTime := time.Now()
		fmt.Fprintf(w, currentTime.Format("15:04"))
	default:
		w.WriteHeader(404)
	}
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		saveText(req.PostForm["author"][0], req.PostForm["entry"][0])
		fmt.Fprintf(w, "Information received: %v\n", req.PostForm)
	default:
		w.WriteHeader(404)
	}
}

func entriesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		saveData, err := os.ReadFile("./miniapi")
		if err == nil {
			fmt.Fprintln(w, string(saveData))
		}
	default:
		w.WriteHeader(404)
	}
}

func saveText(author string, text string) {
	saveFile, err := os.OpenFile("./miniapi", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	defer saveFile.Close()

	w := bufio.NewWriter(saveFile)
	if err == nil {
		fmt.Fprintf(w, "%s:%s\n", author, text)
	}
	w.Flush()
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/entries", entriesHandler)
	http.ListenAndServe(":4567", nil)
}
