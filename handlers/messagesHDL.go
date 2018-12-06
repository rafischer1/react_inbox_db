package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rafischer1/react_inbox_db/models"
)

// Message struct for db query
type Message struct {
	id        int64  `json:"id"`
	read      bool   `json:"read"`
	starred   bool   `json:"starred"`
	selected  bool   `json:"selected"`
	subject   string `json:"subject"`
	body      string `json:"body"`
	labels    string `json:"labels"`
	createdAt string `json:"createdAt"`
	updatedAt string `json:"updatedAt"`
}

func GetAll(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllMessages()
	json.Marshal(data)
	vars := mux.Vars(req)

	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["data"]))
	fmt.Fprintf(w, "data:%s", vars["data"], data)
}

// PostMessage is a function
func PostMessage(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("In the handler post rer", req, w)

	// var Message Message
	// MessageTitle := mux.Vars(req)["title"]
	// data := models.GetAllMessages()

	// if MessageTitle == "" {
	// 	errors.New("user id cannot be empty.")
	// }
	// // this following line must be fixed: cannot assign int to id...
	// // Message.id, _ = strconv.ParseInt(req.FormValue("id"), 0, 32)

	// Message.title = req.FormValue("title")
	// fmt.Println("req Message handler:", Message, data)
	// fmt.Fprint(w, "Content: %v", data)
}

func EditMessage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the handler edit")

	data := models.EditMessage()
	vars := mux.Vars(req)
	fmt.Fprint(w, "Content:", vars["Content"], data)
}

func DeleteMessage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the handler delete")

	data := models.DeleteMessage()
	vars := mux.Vars(req)
	fmt.Fprint(w, "Content:", vars["Content"], data)
}
