package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	vc := VideoCollection{}
	vc.Original = Video{Title: "Hola mundo", URL: "https://example.com/"}
	vc.Original.properties = VideoMeta{Width: 1920, Height: 1080}
	vc.Derivatives = make(map[string]Video)

	awsProcessor := AwsVideo("")
	vc.GeneratePreview(awsProcessor, " - small", VideoMeta{Width: 1920 / 4, Height: 1080 / 4})

	fmt.Printf("Original: (%+v) \n", vc)
}
