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
}
