package strategy

type StrategyType string

const (
	StrategyTypeOAuth    StrategyType = "OAUTH"
	StrategyTypePassword StrategyType = "PASSWORD"
)

type Strategy interface {
	Type() StrategyType
}
