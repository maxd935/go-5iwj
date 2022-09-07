package handler

import (
	"encoding/json"
	"fmt"
	"maxd935/board/structs"
	"net/http"
)

type BoatResponse struct {
	NbrBateau int
}

// GET /boats:
// Devra renvoyer le nombre de bateaux qui sont encore à flot. Si le serveur renvoie 0, cela
// signifie que ce joueur a perdu.

func BoatHandler(shipeSlic []structs.Shipe) func(w http.ResponseWriter, req *http.Request) {
	boatHandlerReturn := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			fmt.Println("GET /boats")
			boatResp := BoatResponse{len(shipeSlic)}
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(boatResp)
		}
	}
	return boatHandlerReturn
}
