package lib

import (
	"errors"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

func RunForm(currencies map[string]string) (float64, string, string) {
	var (
		amount string
		from   string
		to     string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What amount do you want to convert?").
				Value(&amount).
				Validate(func(str string) error {
					_, err := strconv.ParseFloat(str, 64)
					if err != nil {
						return errors.New("invalid amount")
					}
					return nil
				}),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("From what currency do you want to convert?").
				OptionsFunc(func() []huh.Option[string] {
					var currencyOptions []huh.Option[string]
					for currency, name := range currencies {
						currencyOptions = append(currencyOptions, huh.NewOption(name+" ("+currency+")", currency))
					}
					return currencyOptions
				}, &amount).
				Value(&from),
			huh.NewSelect[string]().
				Title("To what currency do you want to convert?").
				OptionsFunc(func() []huh.Option[string] {
					var currencyOptions []huh.Option[string]
					for currency, name := range currencies {
						currencyOptions = append(currencyOptions, huh.NewOption(name+" ("+currency+")", currency))
					}
					return currencyOptions
				}, &amount).
				Value(&to),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	val, _ := strconv.ParseFloat(amount, 64)

	return val, from, to
}
