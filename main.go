package main

import (
	"fmt"

	"github.com/someBroYouKnow/url-shortener-go/form"
)

func main() {
	formData := form.GetFormData()
	err := formData.InitializeForm()

	if err != nil {
		fmt.Println("Error in creating form, try again next time ")
	}
	fmt.Printf("Input URL= %v", formData.InputUrl)

	// Generate some unique non-collisionary ID for the URL

	// Save the ID in redis

	// Setup nginx, or a go lang server that does the hopping from one URL to the other

	// Check if analytics can be placed here

}
