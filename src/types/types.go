// types.go
package types

import "time"

// EventType represents the type of event.
type EventType string

// DispatchType represents the type of alert dispatch.
type DispatchType string

// Event represents an occurrence that may trigger an alert.
type Event struct {
	Client    string
	EventType EventType
	Timestamp time.Time
}

// DispatchStrategy represents the strategy for dispatching alerts.
type DispatchStrategy struct {
	Type    DispatchType
	Message string
	Subject string
}

// AlertConfig represents the configuration for raising alerts.
type AlertConfig struct {
	Client             string
	EventType          EventType
	Type               string
	Count              int
	WindowSizeInSecs   int
	DispatchStrategies []DispatchStrategy
}
