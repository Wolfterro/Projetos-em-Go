package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"external_ip/libs"
)

func printBanner(version float64) {
	fmt.Println("========================")
	fmt.Printf("External IP - Versão %.1f\n", version)
	fmt.Printf("========================\n\n")
}

func printErrorMessage(message string) {
	fmt.Println("Não foi possível obter o endereço IP externo!")
	fmt.Println(message)
}

func getExternalIP() {
	url := "https://ipinfo.io/json"
	response, err := http.Get(url)

	if err != nil {
		printErrorMessage(err.Error())
		os.Exit(1)
	}

	if response.StatusCode == 200 {
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			printErrorMessage(err.Error())
		} else {
			externalIP := new(libs.ExternalIP)

			externalIP.SetJSONData(responseData)
			externalIP.PrintData()
		}
	} else {
		err := fmt.Sprintf("HTTP Status Code: %d", response.StatusCode)
		printErrorMessage(err)
	}
}

func main() {
	version := 1.2

	printBanner(version)
	getExternalIP()
}