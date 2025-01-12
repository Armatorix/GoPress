package config

import (
	"fmt"

	"github.com/Armatorix/GoPress/be/db"
	"github.com/caarlos0/env/v9"
)

type Server struct {
	Port         int  `env:"PORT" envDefault:"8080"`
	RedirectMode bool `env:"REDIRECT_MODE" envDefault:"false"`
}

type Auth struct {
	JwtSecret string `env:"JWT_SECRET,required" envDefault:"2137"`
}

func (s *Server) Address() string {
	return fmt.Sprintf(":%d", s.Port)
}

type Config struct {
	Server Server
	Auth   Auth
	DB     db.Config
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("env parsing: %w", err)
	}
	return cfg, nil
}
