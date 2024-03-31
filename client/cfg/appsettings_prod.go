//go:build prod

package cfg

var AppConfig = map[string]any{
	"server_url": "wss://prodserver.com",
}
