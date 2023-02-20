package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnalagad/docman-api/model"
	"github.com/krishnalagad/docman-api/payload"
)

// controllers
func ServerHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to DocmanMongoDB API by Krishna</h1>"))
}

func CreateDocument(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "POST")

	var document model.Document
	_ = json.NewDecoder(req.Body).Decode(&document)
	data := payload.InsertOneDocument(document)
	json.NewEncoder(res).Encode(data)
	return
}

func GetOneDocument(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(req)
	document := payload.GetOneDocument(params["id"])

	json.NewEncoder(res).Encode(document)
	return
}

func UpdateOneDocument(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(req)
	var document model.Document
	_ = json.NewDecoder(req.Body).Decode(&document)
	result := payload.UpdateOneDocument(document, params["id"])
	json.NewEncoder(res).Encode(result)
	return
}

func GetAllDocuments(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "GET")

	allDocs := payload.GetAllDocuments()
	json.NewEncoder(res).Encode(allDocs)
	return

}

func DeleteOneDocument(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(req)
	count := payload.DeleteOneDocument(params["id"])
	json.NewEncoder(res).Encode("Record deleted with count " + string(rune(count)))
	return
}

func DeleteAllDocuments(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := payload.DeleteAllDocuments()
	json.NewEncoder(res).Encode("Movies deleted successfully " + string(rune(count)))
	return
}
