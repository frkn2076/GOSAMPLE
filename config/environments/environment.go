package environments

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Load() {
	environment := "LOCAL-DEFAULT"
	argument := strings.Join(os.Args, " ")
	if strings.HasPrefix(argument, "go run main.go") && len(os.Args) > 1 {
		environment = os.Args[1]
	} 

	fmt.Println("Environment selected as", environment)
	envFilePath := fmt.Sprintf("config/environments/%s.env", environment)
	godotenv.Load(envFilePath)
}
