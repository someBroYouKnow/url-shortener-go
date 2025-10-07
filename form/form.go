package form

import "github.com/charmbracelet/huh"

type FormData struct {
	InputUrl string
}

var formData = &FormData{
	InputUrl: "",
}

func (f *FormData) InitializeForm() error {

	form := huh.NewForm(
		huh.NewGroup(
			// ask user for a base url
			huh.NewInput().
				Title("Type the website you want a shortener for ").
				Value(&f.InputUrl),
		),
	)

	err := form.Run()
	if err != nil {
		panic(err)
	}

	return nil
}

func GetFormData() *FormData {
	return formData
}
