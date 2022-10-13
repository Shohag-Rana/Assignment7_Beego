package requiredfunction

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Response_Url(url string, cats_data chan string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal()
	}

	cats_data <- string(body)
}
