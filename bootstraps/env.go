package bootstraps

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error while loading the env file")
	}
}
