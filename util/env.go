package util

import (
	"os"

	"github.com/dzakaeryan20/lemonilo/model"
)

func GetEnv() string {
	env := os.Getenv(model.EnvKey)
	if env == "" {
		env = model.LocalEnv
	}
	return env
}
