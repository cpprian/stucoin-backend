package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

type apis struct {
	tasks string
	rewards string
	microCompetencies string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	apis     apis
}

func main() {
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 8000, "HTTP server network port")
	tasksAPI := flag.String("tasksAPI", "http://localhost:4000/tasks", "Tasks API endpoint")
	rewardsAPI := flag.String("rewardsAPI", "http://localhost:4001/rewards", "Rewards API endpoint")
	microCompetenciesAPI := flag.String("microCompetenciesAPI", "http://localhost:4000/micro-competencies", "Micro competencies API endpoint")
	flag.Parse()

	app := &application{
		errorLog: log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(log.Writer(), "INFO\t", log.Ldate|log.Ltime),
		apis: apis{
			tasks: *tasksAPI,
			rewards: *rewardsAPI,
			microCompetencies: *microCompetenciesAPI,
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:     serverURI,
		ErrorLog: app.errorLog,
		Handler:  handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		)(app.routes()),
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Enable CORS");
	app.infoLog.Printf("Starting server on %s", serverURI)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}