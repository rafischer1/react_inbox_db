package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	reqID := req.URL.String()
	id := strings.Split(reqID, "/")[2]

	fmt.Println("In req.URL.String() id:", id)

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
	fmt.Println("In the handler post req.Body:", req.Body)
	body := models.Message{}
	json.NewDecoder(req.Body).Decode(&body)
	fmt.Println("decoder json:", body)

	data := models.PostMessage(body)

	Message := &models.Message{}
	fmt.Println("req Message handler:", Message, data)
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Content: %v", data)
}

// EditMessage handler calls on the model to handle a PUT
func EditMessage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the handler edit")

	data := models.EditMessage()
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Content:", vars["Content"], data)
}

// DeleteMessage sends the delete request to the db
func DeleteMessage(w http.ResponseWriter, req *http.Request) {
	reqID := req.URL.String()
	id := strings.Split(reqID, "/")[2]
	data := models.DeleteMessage(id)
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted Entry:", vars["Deleted Entry"], data, id)
}
