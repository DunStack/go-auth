package auth

type StrategyType string

const (
	StrategyTypeOAuth    StrategyType = "OAUTH"
	StrategyTypePassword StrategyType = "PASSWORD"
)

type StrategiesConfig struct {
	Password StrategyPasswordConfig
	OAuth    StrategyOAuthConfig
}

type StrategyPasswordConfig struct{}

type StrategyOAuthConfig struct{}
