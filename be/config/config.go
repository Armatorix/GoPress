package config

import (
	"fmt"

	"github.com/Armatorix/GoPress/be/api/public"
	"github.com/Armatorix/GoPress/be/db"
	"github.com/Armatorix/GoPress/be/pkg/openai"
	"github.com/caarlos0/env/v9"
)

type Server struct {
	Port         int      `env:"PORT" envDefault:"8080"`
	RedirectMode bool     `env:"REDIRECT_MODE" envDefault:"false"`
	CORSOrigins  []string `env:"CORS_ORIGINS" envDefault:"http://localhost:5173,http://localhost:8080"`
}

type Auth struct {
	JwtSecret string `env:"JWT_SECRET,required" envDefault:"2137"`
}

func (s *Server) Address() string {
	return fmt.Sprintf(":%d", s.Port)
}

type Config struct {
	Server Server
	Public public.Config
	Auth   Auth
	DB     db.Config
	OpenAI openai.Config
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("env parsing: %w", err)
	}
	return cfg, nil
}
