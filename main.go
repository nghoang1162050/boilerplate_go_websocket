package main

import (
	"boilerplate_go_websocket/internal/controller"
	"boilerplate_go_websocket/internal/database"
	"boilerplate_go_websocket/internal/gorm_gen"
	"boilerplate_go_websocket/internal/middleware"
	"boilerplate_go_websocket/internal/router"
	"boilerplate_go_websocket/internal/usecase"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echoMiddleWare "github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, proceeding with environment variables")
    }

	// Initialize database connection
	db, err := database.InitDbClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	query := gorm_gen.Use(db)

	// auth repository, usecase, and controller
	authUseCase := usecase.NewAuthUseCase(query)
	authController := controller.NewAuthController(authUseCase)

	// echo app instance
	e := echo.New()
	e.Use(echoMiddleWare.LoggerWithConfig(echoMiddleWare.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}","host":"${host}","method":"${method}","uri":"${uri}","status":${status}}` + "\n",
		Output: os.Stdout,
	}))

	// Prometheus middleware
    p := prometheus.NewPrometheus("echo", nil)
    p.Use(e)

	apiGroup := e.Group("/api")

	// Authentication middleware
	e.Use(middleware.JWTMiddleware())
	
	router.NewAuthRouter(apiGroup, authController)

	log.Fatal(e.Start(":8080"))

	// hubManager := core.NewHubManager()
	// hubManager.InitHub("default")
	// hubManager.InitHub("secondary")

	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	hubID := r.URL.Query().Get("hubId")
	// 	if hubID == "" {
	// 		hubID = "default"
	// 	}

	// 	hub, ok := hubManager.GetHub(hubID)
	// 	if !ok {
	// 		hub = hubManager.InitHub(hubID)
	// 	}

	// 	upgrader := core.DefaultUpgrader()
	// 	core.ServeWs(hub, w, r, upgrader)
	// })

	// http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
	// 	hubID := r.URL.Query().Get("hubId")
	// 	if hubID == "" {
	// 		http.Error(w, "hub id required", http.StatusBadRequest)
    //         return
	// 	}

	// 	hubManager.CloseHub(hubID)
	// 	w.Write([]byte("Hub closed"))
	// })

	// log.Println("Server started on :8080")
	// http.ListenAndServe(":8080", nil)
}
