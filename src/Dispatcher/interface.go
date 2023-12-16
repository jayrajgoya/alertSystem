package Dispatcher

import (
	"github.com/jayrajgoya/alertSystem/types"
)

type AlertDispatcher interface {
	Dispatch(strategy types.DispatchStrategy)
}
