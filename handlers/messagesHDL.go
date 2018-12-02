package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rafischer1/react_inbox_db/models"
)

type Message struct {
	ID        int      `json:"id"`
	Read      bool     `json:"read"`
	Starred   bool     `json:"starred"`
	Selected  bool     `json:"selected"`
	Subject   string   `sql:"type:varchar(255)"`
	Body      string   `sql:"type:varchar(255)"`
	Labels    []string `sql:",array"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func GetAll(w http.ResponseWriter, req *http.Request) {

	data := models.GetAllMessages()
	vars := mux.Vars(req)

	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["Content"]))
	fmt.Fprintf(w, "Content:%s", vars["Content"], data)
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
