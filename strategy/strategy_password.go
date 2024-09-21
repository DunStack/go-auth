package strategy

var _ Strategy = new(StrategyPassword)

type StrategyPassword struct{}

func (StrategyPassword) Type() StrategyType {
	return StrategyTypePassword
}
