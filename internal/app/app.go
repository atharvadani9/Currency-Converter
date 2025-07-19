package app

import (
	"Currency-Converter/internal/converter"
)

type Application struct {
	converter *converter.Tui
}

func NewApplication() *Application {
	app := &Application{
		converter: &converter.Tui{},
	}
	return app
}

func (a *Application) Run() error {
	return a.converter.CurrencyConverter()
}
