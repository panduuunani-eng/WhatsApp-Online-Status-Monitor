// Improved Error Handling

// Handle connection errors gracefully
func handleConnectionError(err error) {
    if err != nil {
        log.Printf("Connection error: %v", err)
        // Implement retry logic
        retryConnection()
    }
}

// Retry logic for connections
func retryConnection() {
    for i := 0; i < maxRetries; i++ {
        err := connectToWhatsApp()
        if err == nil {
            log.Println("Reconnected successfully")
            return
        }
        log.Printf("Retry %d/%d failed: %v", i+1, maxRetries, err)
        time.Sleep(retryDelay)
    }
    log.Fatal("Failed to reconnect after maximum retries")
}

// Handling presence events
func handlePresenceEvent(event PresenceEvent) {
    // Improved handling of presence events
    switch event.Type {
    case "available":
        // Handle available status
    case "unavailable":
        // Handle unavailable status
    default:
        log.Printf("Unknown presence event type: %s", event.Type)
    }
}

// Context timeout handling
func withTimeout() {
    ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
    defer cancel()

    select {
    case <-ctx.Done():
        log.Println("Operation timed out")
    }
}

// Graceful shutdown support
func shutdownGracefully() {
    // Implementing graceful shutdown logic
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    log.Println("Shutting down gracefully...")
    os.Exit(0)
}