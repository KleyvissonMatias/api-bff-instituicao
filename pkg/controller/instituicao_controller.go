package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var service = ServiceDb{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, request *http.Request) {
	instituicaoEnsino, err := service.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, instituicaoEnsino)
}

func GetByID(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	instituicaoEnsino, err := service.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID da intituição é inválido!")
		return
	}
	respondWithJson(w, http.StatusOK, instituicaoEnsino)
}

func Create(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	var instituicaoEnsino InstituicaoEnsino
	if err := json.NewDecoder(request.Body).Decode(&instituicaoEnsino); err != nil {
		respondWithError(w, http.StatusBadRequest, "Corpo da requisição inválido!")
		return
	}
	instituicaoEnsino.ID = bson.NewObjectId()
	if err := service.Create(instituicaoEnsino); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, instituicaoEnsino)
}

func Update(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	params := mux.Vars(request)
	var instituicaoEnsino InstituicaoEnsino
	if err := json.NewDecoder(request.Body).Decode(&instituicaoEnsino); err != nil {
		respondWithError(w, http.StatusBadRequest, "Corpo da requisição inválido!")
		return
	}
	if err := service.Update(params["id"], instituicaoEnsino); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result:": instituicaoEnsino.Nome + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	params := mux.Vars(request)
	if err := service.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
