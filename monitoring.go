package main

import (
	"fmt"
	"sync"
	"time"
)

// EventProcessor defines the interface for processing events.
type EventProcessor interface {
	ProcessEvent(event Event)
}

// MonitoringService monitors events and raises alerts based on configurations.
type MonitoringService struct {
	AlertConfigList []AlertConfig
	Events          []Event
	mu              sync.Mutex
	AlertingService AlertingService
}

// NewMonitoringService creates a new instance of MonitoringService.
func NewMonitoringService(alertConfigList []AlertConfig, alertingService AlertingService) *MonitoringService {
	return &MonitoringService{
		AlertConfigList: alertConfigList,
		AlertingService: alertingService,
	}
}

// StartMonitoring starts monitoring events for a given client and event type.
func (ms *MonitoringService) StartMonitoring(client string, eventType EventType) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	// Current time
	currentTime := time.Now()

	// Time threshold for sliding window (1 hour)
	windowThreshold := currentTime.Add(-time.Hour)

	// Filter events within the last hour
	var eventsWithinWindow []Event
	for _, event := range ms.Events {
		if event.Timestamp.After(windowThreshold) {
			eventsWithinWindow = append(eventsWithinWindow, event)
		}
	}

	for _, config := range ms.AlertConfigList {
		if config.Client == client && config.EventType == eventType {
			fmt.Printf("[INFO] MonitoringService: Client %s %s %s starts\n", client, eventType, config.Type)

			// Simulate events
			for i := 1; i <= config.Count; i++ {
				ms.Events = append(ms.Events, Event{
					Client:    client,
					EventType: eventType,
					Timestamp: currentTime,
				})
			}

			fmt.Printf("[INFO] MonitoringService: Client %s %s %s ends\n", client, eventType, config.Type)

			// Check if the alert threshold is breached within the last hour
			if len(eventsWithinWindow) >= config.Count {
				fmt.Printf("[INFO] MonitoringService: Client %s %s %s threshold breached\n", client, eventType, config.Type)

				// Dispatch alerts
				for _, strategy := range config.DispatchStrategies {
					ms.AlertingService.DispatchAlert(strategy)
				}

				// Clear events outside the sliding window
				var newEvents []Event
				for _, event := range ms.Events {
					if event.Timestamp.After(windowThreshold) {
						newEvents = append(newEvents, event)
					}
				}
				ms.Events = newEvents
			}
			return
		}
	}
}

// ProcessEvent processes an incoming event.
func (ms *MonitoringService) ProcessEvent(event Event) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	ms.Events = append(ms.Events, event)
	ms.StartMonitoring(event.Client, event.EventType)
}
