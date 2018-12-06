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
func GetAllFiles(w http.ResponseWriter, req *http.Request) {
	fmt.Println("in the getall files handler")
	data := models.GetAllFilesMDL()

	//return the data
	fmt.Printf("data: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)

}

// GetOne handler to handle one record
func GetOneFile(w http.ResponseWriter, req *http.Request) {

	reqID := req.URL.String()
	id := strings.Split(reqID, "/")[2]

	fmt.Println("In req.URL.String() id:", id)

	data := models.GetOneFileMDL(id)
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
func PostFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the handler post req.Body:", req.Body)
	body := models.File{}

	json.NewDecoder(req.Body).Decode(&body)
	fmt.Println("decoder json:", body)

	// gonna need to parse these to look like this: 2018-12-06 11:35:13

	data := models.PostFileMDL(body)

	File := &models.File{}
	fmt.Println("req File handler:", File, data)
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Content: %v", data)
}

// EditMessage handler calls on the model to handle a PUT
func EditFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the handler edit")

	data := models.EditFileMDL()
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Content:", vars["Content"], data)
}

// DeleteMessage sends the delete request to the db
func DeleteFile(w http.ResponseWriter, req *http.Request) {
	reqID := req.URL.String()
	id := strings.Split(reqID, "/")[2]
	data := models.DeleteFileMDL(id)
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted File:", vars["Deleted File"], data, id)
}
