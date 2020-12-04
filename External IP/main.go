package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"external_ip/libs"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	version := 1.1

	printBanner(version)
	getExternalIP()
}

func printBanner(version float64) {
	fmt.Println("========================")
	fmt.Printf("External IP - Versão %.1f\n", version)
	fmt.Printf("========================\n\n")
}

func getExternalIP() {
	url := "https://ipinfo.io/json"
	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Não foi possível obter o endereço IP externo!")
		fmt.Println(err)

		os.Exit(1)
	}

	if response.StatusCode == 200 {
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("Não foi possível obter o endereço IP externo!")
			fmt.Println(err)
		} else {
			externalIP := new(libs.ExternalIP)

			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			json.Unmarshal(responseData, externalIP)

			externalIP.PrintData()
		}
	} else {
		fmt.Println("Não foi possível obter o endereço IP externo!")
		fmt.Println("HTTP Status Code:", response.StatusCode)
	}
}
