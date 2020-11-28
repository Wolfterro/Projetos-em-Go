package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	version := 1.0
	printBanner(version)

	getExternalIP()
}

func printBanner(version float64) {
	fmt.Println("========================")
	fmt.Printf("External IP - Versão %.1f\n", version)
	fmt.Println("========================")
}

func getExternalIP() {
	url := "https://ipinfo.io/json"
	response, _ := http.Get(url)

	if response.StatusCode == 200 {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			responseString := string(responseData)
			fmt.Println(responseString)
		}
	} else {
		fmt.Println("Não foi possível obter o endereço IP externo!")
	}
}