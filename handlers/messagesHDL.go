package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/rafischer1/react_inbox_db/models"
	m "github.com/rafischer1/react_inbox_db/models"
)

// GetAll handler to handle all records
func GetAll(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("in the getall handler", req)
	data := m.GetAllMessages()

	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)

}

// GetOne handler to handle one record
func GetOne(w http.ResponseWriter, req *http.Request) {

	reqID := req.URL.String()
	id := strings.Split(reqID, "/")[2]

	fmt.Println("In req.URL.String() id:", id)

	data := m.GetOneMessage(id)
	json.Marshal(data)

	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)
}

// PostMessage is a function
func PostMessage(w http.ResponseWriter, req *http.Request) {

	fmt.Println("In the handler post req.Body:", req.Body)
	body := m.Message{}
	fmt.Println("before new encoder", body)

	json.NewDecoder(req.Body).Decode(body)
	body.Subject = "hi there"
	body.Read = true
	body.Selected = true
	body.Starred = false
	body.Body = "Test one"
	body.Labels = "{'personal', 'dev'}"
	fmt.Println("decoder json:", body.Subject)

	data, err := models.PostMessage(body.Read, body.Starred, body.Selected, body.Subject, body.Body, body.Labels, body.CreatedAt, body.UpdatedAt)
	if err != nil {
		panic(err)
	}

	Message := &models.Message{}
	fmt.Println("req Message handler:", Message, data)
	w.Header().Set("Content-Type", "application/json")
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
	// vars := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted Entry:", data)
}
