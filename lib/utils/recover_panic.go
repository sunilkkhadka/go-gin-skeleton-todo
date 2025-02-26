package utils

import (
	"boilerplate-api/lib/config"
)

// RecoverPanic recovers panic in the application
func RecoverPanic(logger config.Logger) func() {
	return func() {
		if r := recover(); r != nil {
			logger.Info("☠️ panic recovered: ", r)
		}
	}
}
