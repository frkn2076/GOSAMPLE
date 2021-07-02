package environments

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	environment := "LOCAL-DEFAULT"
	if len(os.Args) > 1 {
		environment = os.Args[1]
	} 

	fmt.Println("Environment selected as", environment)
	envFilePath := fmt.Sprintf("config/environments/%s.env", environment)
	godotenv.Load(envFilePath)
}
