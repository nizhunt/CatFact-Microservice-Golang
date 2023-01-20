package main

import (
	"context"
	"encoding/json"
	"net/http"
)

// rewrite this using go-fiber
// holds interface, can make new service without touching the old one
type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr string) error {
	// routes with handlers
	http.HandleFunc("/", s.handleGetCatFact)
	// serve to the given address
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		writeJson(w, http.StatusUnprocessableEntity, map[string]any{
			"error": err.Error(),
		})
	}
	writeJson(w, http.StatusOK, fact)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
