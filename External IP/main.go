package main

import (
	"fmt"
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
	data, err := libs.GetIPInfoData()

	if err != nil {
		printErrorMessage(err.Error())
		os.Exit(1)
	}

	externalIP := new(libs.ExternalIP)

	externalIP.SetJSONData(data)
	externalIP.PrintData()
}

func main() {
	version := 1.3

	printBanner(version)
	getExternalIP()
}