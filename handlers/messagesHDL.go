package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rafischer1/react_inbox_db/models"
)

// GetAll handler to handle all records
func GetAll(w http.ResponseWriter, req *http.Request) {
	fmt.Println("in the getall handler", req)
	data := models.GetAllMessages()
	json.Marshal(data)
	vars := mux.Vars(req)

	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["data"]))
	fmt.Fprintf(w, "data:%s", vars["data"], data)

}

// GetOne handler to handle one record
func GetOne(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the getOne handler:", req)
	id := req.FormValue("id")
	data := models.GetOneMessage(id)
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
	data := models.GetAllMessages()
	Message := &models.Message{}

	// MessageSubject := mux.Vars(req)["subject"]

	// if MessageTitle == "" {
	// 	errors.New("user id cannot be empty.")
	// }
	// this following line must be fixed: cannot assign int to id...
	// Message.id, _ = strconv.ParseInt(req.FormValue("id"), 0, 32)

	// Message.subject = req.FormValue("subject")
	fmt.Println("req Message handler:", Message, data)
	fmt.Fprint(w, "Content: %v", data)
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
