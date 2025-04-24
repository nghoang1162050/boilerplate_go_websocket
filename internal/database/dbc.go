package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDbClient() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass	:= os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	// metricPort, _ := strconv.Atoi(os.Getenv("HTTP_METRIC_PORT"))

	db, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}
	
	// db.Use(prometheus.New(prometheus.Config{
	// 	DBName: dbName,
	// 	RefreshInterval: 15,
	// 	PushAddr: dbHost, // push metrics if `PushAddr` configured
	// 	StartServer: true,
	// 	HTTPServerPort:  uint32(metricPort),  // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
	// 	MetricsCollector: []prometheus.MetricsCollector {
	// 	  &prometheus.MySQL{
	// 		VariableNames: []string{"Threads_running"},
	// 	  },
	// 	},  // user defined metrics
	//   }))

	return db, err
}