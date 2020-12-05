package main

import (
	"fmt"

	"Playah/libs"
)

// VERSION of the program
const VERSION = "1.0.0"

func printProgramBanner() {
	fmt.Println("=========================")
	fmt.Printf("# Playah - Versão %s #\n", VERSION)
	fmt.Printf("=========================\n")
}

func printFileDetail(mp3 *libs.MP3Content) {
	fmt.Printf("♫ ♪ Tocando '%s' ♪ ♫ \n", mp3.Filename)
}

func main() {
	printProgramBanner()
	path, sampleRate := libs.ParseArgs()

	mp3 := new(libs.MP3Content)
	mp3.BuildMP3Player(path, sampleRate)

	printFileDetail(mp3)
	mp3.Play()
}
