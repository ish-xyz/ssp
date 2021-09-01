package ui

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ish-xyz/ssp/internal/logger"
)

func backendGet(w http.ResponseWriter, r *http.Request, e string) (*Request, error) {
	var req *Request

	logger.DebugLogger.Printf("Frontend calling backend at endpoint %s", e)

	reqURL := fmt.Sprintf("http://%s/%s", appConfig.BackendAddr, e)
	resp, err := http.Get(reqURL)
	if err != nil {
		logger.ErrorLogger.Printf("Can't request backend %s. Error => %v ", e, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&req)

	return req, nil
}

func backendPost(w http.ResponseWriter, r *http.Request, e string, d interface{}) (*Request, error) {
	fmt.Println(w, r, e, d)
	return nil, nil
}
