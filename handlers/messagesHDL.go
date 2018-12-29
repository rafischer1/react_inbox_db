package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/rafischer1/react_inbox_db/models"
	m "github.com/rafischer1/react_inbox_db/models"
)

type Reader interface {
	Read(buf []byte) (n int, err error)
}

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
	fmt.Printf("In the handler post req.Body: %+v", req.Method)
	if req.Method == "OPTIONS" {
		fmt.Println("Options in POST:", req.Method)
	}
	if req.Method == "POST" {
		var bodyBytes []byte
		if req.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(req.Body)
		}

		// Restore the io.ReadCloser to its original state
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)
		// body := m.Message{}
		str := bodyString
		res := m.Message{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println("Res subject:", res.Subject, "res body:", res.Body)
		// json.NewDecoder(req.Body).Decode(body)
		// there is a problem here with ID - maybe have to solve that on the model side with a query to determine last recorded ID although I don't understadn why they don't incremenet

		data, err := models.PostMessage(res.Subject, res.Body)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Content: %v", data)
	}
}

// EditMessage handler calls on the model to handle a PUT
func EditMessage(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	if req.Method == "OPTIONS" {
		fmt.Println("Options in EDIT:", req.Method)
	}
	if req.Method == "PUT" {
		fmt.Println("In the handler edit method, req.id:", req.Method, req.Body)
		// something not right with the res and req.Body but setup is ok
		var bodyBytes []byte
		if req.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(req.Body)
		}

		// Restore the io.ReadCloser to its original state
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)
		// body := m.Message{}
		str := bodyString
		res := m.Message{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println("Res", res, "res body:", res.Body)
		// json.NewDecoder(req.Body).Decode(body)
		// there is a problem here with ID - maybe have to solve that on the model side with a query to determine last recorded ID although I don't understadn why they don't incremenet

		data, err := models.EditMessage(res.ID, res.Body)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Content: %v", data)
	}
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
