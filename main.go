package main

import (
	"DailyTasks/Database"
	"DailyTasks/Progress"
	"DailyTasks/Tasks"
	"DailyTasks/config"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := Database.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
	}()

	// TODO Create repositories
	// TODO Create Services
	// TODO Create Handlers

	r := createRouter()
	srv := createServer(cfg, r)

	stopChan, errChan := createShutdownChannels()

	// Start the server in a goroutine
	go func() {
		log.Printf("Server starting on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("failed to start server: %w", err)
		}
	}()

	waitForSignals(errChan, stopChan, srv)
}

func createRouter() *chi.Mux {
	r := chi.NewRouter()
	setupMiddleware(r)
	defineRoutes(r)
	return r
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)                    // Consider structured logging
	r.Use(middleware.Recoverer)                 // Recover from panics
	r.Use(middleware.Timeout(60 * time.Second)) // Set request timeout

	// TODO: Add CORS middleware if your React app is on a different domain/port
	// e.g., using github.com/go-chi/cors
	// corsMiddleware := cors.New(cors.Options{ ... })
	// r.Use(corsMiddleware.Handler)
}

func defineRoutes(r *chi.Mux) {
	// Example using a placeholder function from an internal / api package:
	// api.RegisterRoutes(r, userHandler) // Pass the router and necessary handlers

	// Placeholder route
	r.Get("/AllTasks", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		_, err := fmt.Fprintln(w, "Server is running!")
		if err != nil {
			log.Printf("Failed to write response: %v", err)
			return
		}
	})
}

func createServer(cfg *config.Config, r *chi.Mux) *http.Server {
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return srv
}

func createShutdownChannels() (chan os.Signal, chan error) {
	// Channel to listen for interrupt or terminate signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// Channel to listen for server errors
	errChan := make(chan error, 1)
	return stopChan, errChan
}

func waitForSignals(errChan chan error, stopChan chan os.Signal, srv *http.Server) {
	select {
	case err := <-errChan:
		log.Fatalf("Server error: %v", err)
	case sig := <-stopChan:
		log.Printf("Received signal: %v. Shutting down gracefully...", sig)

		// Create a context with a timeout for shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Attempt graceful shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown failed: %v", err)
		} else {
			log.Println("Server stopped gracefully.")
		}
	}
}

func testInMemoryRepositories() {
	taskRepository := Tasks.Repository(Tasks.NewMemoryRepository())
	progressRepository := Progress.Repository(Progress.NewMemoryRepository())

	err := fillRepositories(taskRepository, progressRepository)
	if err != nil {
		fmt.Println(err)
		return
	}

	allTasks, err := taskRepository.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	tasksWithProgress := make(map[Tasks.Task]Progress.PrintableProgress)
	for _, task := range allTasks {
		printableProgress, found, err := progressRepository.GetAllProgress(task.Uuid)
		if err != nil {
			fmt.Printf("Error while getting progress for task %s:\n\n%v", task.Uuid, err)
			continue
		}
		if !found {
			fmt.Printf("Progress not found for task %s\n", task.Uuid)
			continue
		}
		tasksWithProgress[task] = printableProgress
	}

	printProgress(tasksWithProgress)
}

func fillRepositories(taskRepository Tasks.Repository, progressRepository Progress.Repository) error {
	booleanTask := Tasks.NewTask(Tasks.BooleanTask, "booleanTask")
	numberTask := Tasks.NewTask(Tasks.NumberTask, "numberTask")
	durationTask := Tasks.NewTask(Tasks.DurationTask, "durationTask")

	err := taskRepository.Add(booleanTask)
	if err != nil {
		return err
	}
	err = taskRepository.Add(numberTask)
	if err != nil {
		return err
	}
	err = taskRepository.Add(durationTask)
	if err != nil {
		return err
	}

	err = progressRepository.AddTask(booleanTask)
	if err != nil {
		return err
	}
	err = progressRepository.AddTask(numberTask)
	if err != nil {
		return err
	}
	err = progressRepository.AddTask(durationTask)
	if err != nil {
		return err
	}

	err = progressRepository.UpdateBooleanProgress(booleanTask.Uuid, today(), true)
	if err != nil {
		return err
	}
	err = progressRepository.UpdateNumberProgress(numberTask.Uuid, today(), 42)
	if err != nil {
		return err
	}
	err = progressRepository.UpdateDurationProgress(durationTask.Uuid, today(), time.Minute)
	if err != nil {
		return err
	}

	return nil
}
