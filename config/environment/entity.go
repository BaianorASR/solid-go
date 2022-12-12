package config

// Environment is a struct that contains all the environment variables
type Environment struct {
	Port             int    `env:"PORT"`
	Env              string `env:"ENV"`
	CaseSensitiveURL bool   `env:"CASE_SENSITIVE_URL"`
}
