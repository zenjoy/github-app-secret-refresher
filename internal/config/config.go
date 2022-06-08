package config

import (
	"log"
	"os"
	"strconv"

	"github.com/disturbing/github-app-k8s-secret-refresher/v2/internal/types"
	"github.com/joho/godotenv"
)

var (
	TokenProcessorType      types.TokenProcessorType
	GithubAppId             int
	GithubAppInstallationId int
	GithubAppPrivateKeyFile string

	KubeConfigPath                       string
	KubeSecretName                       string
	KubeSecretAuthUsernameKey            string
	KubeSecretInstallationAccessTokenKey string
	KubeSecretNamespace                  string
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Load() {
	godotenv.Load()

	TokenProcessorType = types.TokenProcessorType(getEnv("TOKEN_PROCESSOR_TYPE", "KUBERNETES"))
	GithubAppId = getEnvAsInt("GITHUB_APP_ID")
	GithubAppInstallationId = getEnvAsInt("GITHUB_APP_INSTALLATION_ID")
	GithubAppPrivateKeyFile = os.Getenv("GITHUB_APP_PRIVATE_KEY_PATH")

	KubeConfigPath = os.Getenv("KUBE_CONFIG_PATH")
	KubeSecretName = getEnv("KUBE_SECRET_NAME", "github-credentials")
	KubeSecretAuthUsernameKey = getEnv("KUBE_SECRET_AUTH_USERNAME_KEY", "username")
	KubeSecretInstallationAccessTokenKey = getEnv("KUBE_SECRET_INSTALLATION_ACCESS_TOKEN_KEY", "password")
	KubeSecretNamespace = os.Getenv("KUBE_SECRET_NAMESPACE")
}

func getEnvAsInt(envVar string) int {
	if val := os.Getenv(envVar); val != "" {
		intVal, err := strconv.Atoi(val)

		if err == nil {
			return intVal
		}

		log.Panicf("Environment variable %s is not an int", envVar)
	}

	log.Panicf("Environment variable %s is not an int", envVar)
	return 0
}
