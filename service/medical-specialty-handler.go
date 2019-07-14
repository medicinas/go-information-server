package service

import (
	"encoding/json"
	"medicinas/information-server/models"
	"net/http"
	"strconv"
)

func GetSomeResponse(w http.ResponseWriter, r *http.Request) {
	specialty := models.MedicalSpecialty{}
	data, _ := json.Marshal(specialty)
	writeJsonResponse(w, http.StatusOK, data)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
