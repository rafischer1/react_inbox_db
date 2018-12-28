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
	enableCors(&w)
	// fmt.Println("in the getall handler", req)
	data := m.GetAllMessages()

	//return the data
	// fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Hit the getAll messages route:", http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)

}

// GetOne handler to handle one record
func GetOne(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
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
	enableCors(&w)
	fmt.Printf("In the handler post req.Body: %v", req.Method)
	if req.Method == "OPTIONS" {
		fmt.Println("Options in POST")
	}
	if req.Method == "POST" {
		body := m.Message{}

		fmt.Println("before new encoder", &req.Body)

		json.NewDecoder(req.Body).Decode(body)
		// there is a problem here with ID - maybe have to solve that on the model side with a query to determine last recorded ID although I don't understadn why they don't incremenet

		postID, postSubject, err := models.PostMessage(body.Subject, body.Body)
		if err != nil {
			panic(err)
		}

		Message := &models.Message{}
		fmt.Println("req Message handler:", Message)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, "Content: %v", postID, postSubject)
	}

}

// EditMessage handler calls on the model to handle a PUT
func EditMessage(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	fmt.Println("In the handler edit")

	data := models.EditMessage()
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Content:", vars["Content"], data)
}

// DeleteMessage sends the delete request to the db
func DeleteMessage(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	reqID := req.URL.String()
	id := strings.Split(reqID, "/")[2]
	data := models.DeleteMessage(id)
	// vars := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted Entry:", data)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}
