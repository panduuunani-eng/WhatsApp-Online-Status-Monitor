package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Set up signal handling for graceful shutdown
    go func() {
        signals := make(chan os.Signal, 1)
        signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
        <-signals
        log.Println("Received shutdown signal")
        cancel() // Cancel the context on signal
    }()

    // Create an HTTP server
    srv := &http.Server{
        Addr: ":8080",
        Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("Hello, world!"))
        }),
    }

    // Start the server in a goroutine
    go func() {
        log.Println("Starting server on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe(): %v", err)
        }
    }() 

    // Wait for the context to be cancelled
    <-ctx.Done()
    log.Println("Shutting down server...")

    // Set a timeout for the shutdown
    ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelShutdown()

    if err := srv.Shutdown(ctxShutdown); err != nil {
        log.Fatalf("Server Shutdown Failed:%+v", err)
    }
    log.Println("Server exited properly")
}