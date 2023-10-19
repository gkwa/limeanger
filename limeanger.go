package limeanger

import (
	"log/slog"
)

func Main() int {
	slog.Debug("limeanger", "test", true)
	test()
	return 0
}
