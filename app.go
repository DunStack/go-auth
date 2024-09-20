package auth

func NewApp(cfg *Config) *App {
	return &App{
		cfg: cfg,
	}
}

type App struct {
	cfg *Config
}

func (a *App) NewGraphQLHandler(cfg *GraphQLConfig) *GraphQLHandler {
	return newGraphQLHandler(cfg)
}
