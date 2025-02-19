package app

import (
	"log/slog"

	"github.com/kyuff/anchor"
)

var logger = anchor.Singleton(func() (*slog.Logger, error) {
	return slog.Default(), nil
})
