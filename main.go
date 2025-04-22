package main

import (
	"boilerplate_go_websocket/internal/dto"
	"log"
	"net/http"
)

func main() {
	hubManager := dto.NewHubManager()
	hubManager.InitHub("default")
	hubManager.InitHub("secondary")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hubID := r.URL.Query().Get("hubId")
		if hubID == "" {
			hubID = "default"
		}

		hub, ok := hubManager.GetHub(hubID)
		if !ok {
			hub = hubManager.InitHub(hubID)
		}

		upgrader := dto.DefaultUpgrader()
		dto.ServeWs(hub, w, r, upgrader)
	})

	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		hubID := r.URL.Query().Get("hubId")
		if hubID == "" {
			http.Error(w, "hub id required", http.StatusBadRequest)
            return
		}

		hubManager.CloseHub(hubID)
		w.Write([]byte("Hub closed"))
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
