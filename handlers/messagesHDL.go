package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rafischer1/react_inbox_db/models"
)

type Message struct {
	ID        int      `json:"ID"`
	Read      bool     `json:"Read"`
	Starred   bool     `json:"Starred"`
	Selected  bool     `json:"Selected"`
	Subject   string   `json:"Subject"`
	Body      string   `json:"Body"`
	Labels    []string `sql:",array"`
	CreatedAt string   `json:"CreatedAt"`
	UpdatedAt string   `json:"UpdatedAt"`
}

func GetAll(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllMessages()
	json.Marshal(data)
	vars := mux.Vars(req)

	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["data"]))
	fmt.Fprintf(w, "data:%s", vars["data"], &data)
}

//Rest of REST routes
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
