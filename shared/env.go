package shared

import (
	"fmt"
	"os"
	"strconv"
)

var (
	EnvModeDevelopment = "development"
	EnvModeProduction  = "production"
)

var (
	EnvEnvironment = env("ENVIRONMENT_MODE", EnvModeDevelopment)
	EnvServerPort  = envInt("SERVER_PORT", 7172)
	EnvBindAddress = env("SERVER_IP_BIND", "")
)

var (
	envs = make(map[string]string)
)

// env gives you either value of os.Getenv or the default value specified
func env(name string, defaultValue string) string {
	envs[name] = defaultValue

	variable := os.Getenv(name)

	if variable == "" {
		return defaultValue
	}

	return variable
}

func envInt(name string, defaultValue int) int {
	envVar := env(name, strconv.Itoa(defaultValue))
	result, err := strconv.Atoi(envVar)
	if err != nil {
		return defaultValue
	}
	return result
}

func printEnvFile() {
	fmt.Println("=== EXAMPLE ENV FILE ===")
	for k, v := range envs {
		fmt.Printf("%s=%s\n", k, v)
	}
	fmt.Println("=== END OF EXAMPLE ENV FILE ===")
}
