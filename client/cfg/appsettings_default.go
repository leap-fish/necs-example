//go:build !dev && !prod

package cfg

var AppConfig = map[string]any{
	"server_url": "ws://localhost:7172",
}
