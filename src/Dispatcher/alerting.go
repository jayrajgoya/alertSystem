package Dispatcher

import "fmt"

// ConsoleAlertDispatcher implements AlertDispatcher for console alerts.
type ConsoleAlertDispatcher struct{}

func (c ConsoleAlertDispatcher) Dispatch(strategy DispatchStrategy) {
	fmt.Printf("[INFO] AlertingService: Dispatching to Console\n")
	fmt.Printf("[WARN] Alert: `%s`\n", strategy.Message)
}

// EmailAlertDispatcher implements AlertDispatcher for email alerts.
type EmailAlertDispatcher struct{}

func (e EmailAlertDispatcher) Dispatch(strategy DispatchStrategy) {
	fmt.Printf("[INFO] AlertingService: Dispatching an Email\n")
	// Implement email dispatch logic here
	// You can use a third-party library or a service for sending emails
}

// AlertingService orchestrates the alert dispatching process.
type AlertingService struct {
	Dispatchers map[DispatchType]AlertDispatcher
}

// NewAlertingService creates a new instance of AlertingService.
func NewAlertingService() *AlertingService {
	return &AlertingService{
		Dispatchers: make(map[DispatchType]AlertDispatcher),
	}
}

// AddDispatcher adds a new dispatcher to the alerting service.
func (as *AlertingService) AddDispatcher(dispatchType DispatchType, dispatcher AlertDispatcher) {
	as.Dispatchers[dispatchType] = dispatcher
}

// DispatchAlert dispatches alerts based on the specified strategy.
func (as *AlertingService) DispatchAlert(strategy DispatchStrategy) {
	dispatcher, exists := as.Dispatchers[strategy.Type]
	if !exists {
		fmt.Printf("[WARN] Unknown dispatch type: %s\n", strategy.Type)
		return
	}
	dispatcher.Dispatch(strategy)
}
