package kitchenservice

// func KitchenService() {
//     // Set up RabbitMQ client once
//     client, err := NewClient()
//     if err != nil {
//         log.Fatalf("Error initializing RabbitMQ client: %v", err)
//     }

//     // Set up graceful shutdown
//     sigs := make(chan os.Signal, 1)
//     signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
//     done := make(chan bool, 1)

//     // Start a goroutine to handle messages
//     go func() {
//         msgs, err := client.Consume("kitchen_queue")
//         if err != nil {
//             log.Fatalf("Error consuming messages: %v", err)
//         }

//         for msg := range msgs {
//             log.Printf("Received message: %s", msg.Body)
//             // Process the message...
//             msg.Ack(false)
//         }
//     }()

//     // Block until a signal is received
//     <-sigs
//     log.Println("Shutting down gracefully...")
//     client.Close()
//     done <- true
//     <-done
//     log.Println("Service stopped")
// }