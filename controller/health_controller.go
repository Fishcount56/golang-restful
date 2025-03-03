package controller

import (
	"encoding/json"
	"net/http"
	"restful-api/dto/response"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func HealthController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res := response.NewApiResponseBuilder().SetStatus(200).SetMessage("OK").SetRequestId(uuid.New().String()).Build()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
