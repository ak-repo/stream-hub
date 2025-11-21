package helper

import (
	"time"

	"github.com/ak-repo/stream-hub/config"
)

// overrideLocal ensures localhost usage in local development mode
func OverrideLocal(cfg *config.Config) {
	if cfg.App.Environment != "development" {
		return
	}

	cfg.Services.Gateway.Host = "localhost"
	cfg.Services.Auth.Host = "localhost"
	cfg.Services.Chat.Host = "localhost"
	cfg.Services.File.Host = "localhost"
	cfg.Services.Notification.Host = "localhost"

}

func TimeToString(t time.Time) string {
	const layout = "2006-01-02 15:04:05"
	return t.Format(layout)
}
