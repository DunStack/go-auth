package strategy

type StrategyType string

const (
	StrategyTypeOAuth StrategyType = "OAUTH"
)

type Strategy interface {
	Type() StrategyType
}
