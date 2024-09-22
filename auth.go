package auth

type Config struct {
	DB         *DBConfig
	Token      *TokenConfig
	Strategies *StrategiesConfig
}
