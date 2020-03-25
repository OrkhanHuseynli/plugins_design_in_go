package controller

import (
	"encoding/json"
	"net/http"
	"plugins_design_in_go/src/models"
	"plugins_design_in_go/src/plugins/dbplugin"
)

type SimpleHandler struct {
	dbPlugin *dbplugin.DbPlugin
}

func NewSimpleHandler(dbPlugin *dbplugin.DbPlugin) http.Handler {
	return SimpleHandler{dbPlugin:dbPlugin}
}

func (h SimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.handlePostRequest(w, r)
	default :
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h SimpleHandler) handlePostRequest(w http.ResponseWriter, r *http.Request){
	var payment models.Payment
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payment)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if payment.Author == "" ||  payment.Sum == "" || payment.Product == "" {
		http.Error(w, "Required body fields are empty", http.StatusUnprocessableEntity)
		return
	}
	h.dbPlugin.Repository.AddPayment(&payment)
	response := models.SimpleResponse{Message: "Your payment was successfully put in the process" }
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
