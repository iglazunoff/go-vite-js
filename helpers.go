package go_vite_js

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const logTitle = "> ViteJS"

func checkIsProduction(viteDevClient string) bool {
	client := http.DefaultClient
	client.Timeout = time.Second

	response, err := client.Get(viteDevClient)

	if err != nil || response.StatusCode >= 500 {
		info("ViteJS run in PRODUCTION mode")

		return true
	}

	info("ViteJS run in DEVELOPMENT mode")

	return false
}

func makeUriForDevClient(developmentUri string) string {
	return developmentUri + "/" + defaultDevServerClient
}

func importManifestItems(path string) *ManifestItems {
	jsonData, err := os.ReadFile(path)

	if err != nil {
		info(fmt.Sprintf("Failed read manifest.json, reason: %v", err.Error()))
	}

	info(fmt.Sprintf("Inspected manifest.json by path: %v", path))

	manifestItems := make(ManifestItems)

	err = json.Unmarshal(jsonData, &manifestItems)

	if err != nil {
		fail(fmt.Sprintf("Parsing of manifest.json failed with error: %v", err.Error()))
	}

	return &manifestItems
}

func isValidHttpPortOrFail(httpPort int) {
	if httpPort < 1 || httpPort > 65535 {
		panic(fmt.Sprintf("Given invalid http port '%v'", httpPort))
	}
}

func isValidHttpSchemeOrFail(httpScheme string) {
	if httpScheme != "http" && httpScheme != "https" {
		panic(fmt.Sprintf("Given invalid http scheme '%v'", httpScheme))
	}
}

func info(message string) {
	log.Printf("%v | %v", logTitle, message)

	line := ""
	for i := 0; i < 120; i++ {
		line += "-"
	}

	log.Printf("%v", line)
}

func fail(message string) {
	log.Fatalf("%v | %v", logTitle, message)
}
