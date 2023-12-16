// main.go
package main

import "time"

func main() {
	// Sample configuration
	alertConfigList := []AlertConfig{
		{
			Client:           "X",
			EventType:        "PAYMENT_EXCEPTION",
			Type:             "TUMBLING_WINDOW",
			Count:            10,
			WindowSizeInSecs: 10,
			DispatchStrategies: []DispatchStrategy{
				{
					Type:    "CONSOLE",
					Message: "issue in payment",
				},
				{
					Type:    "EMAIL",
					Subject: "payment exception threshold breached",
				},
			},
		},
		{
			Client:           "X",
			EventType:        "USERSERVICE_EXCEPTION",
			Type:             "SLIDING_WINDOW",
			Count:            10,
			WindowSizeInSecs: 10,
			DispatchStrategies: []DispatchStrategy{
				{
					Type:    "CONSOLE",
					Message: "issue in user service",
				},
			},
		},
	}

	// Initialize alerting service with dispatchers
	alertingService := NewAlertingService()
	alertingService.AddDispatcher("CONSOLE", ConsoleAlertDispatcher{})
	alertingService.AddDispatcher("EMAIL", EmailAlertDispatcher{})

	// Initialize monitoring service with the configuration and alerting service
	monitoringService := NewMonitoringService(alertConfigList, *alertingService)

	// Simulate events to trigger alerts
	event := Event{
		Client:    "X",
		EventType: "PAYMENT_EXCEPTION",
		Timestamp: time.Now(),
	}
	monitoringService.ProcessEvent(event)

	// Keep the main goroutine running
	select {}
}
