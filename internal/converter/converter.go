package converter

import (
	"fmt"
	"github.com/charmbracelet/huh"
)

type Tui struct {
	FromCurrency string
	ToCurrency   string
	Amount       string
	Result       string
}

func (t *Tui) CurrencyConverter() error {
	// Set From Currency to always be USD
	t.FromCurrency = "USD"

	currencies := []huh.Option[string]{
		huh.NewOption("Euro", "EUR"),
		huh.NewOption("British Pound", "GBP"),
		huh.NewOption("Japanese Yen", "JPY"),
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("To Currency").
				Options(currencies...).
				Value(&t.ToCurrency),

			huh.NewInput().
				Title("Amount (USD)").
				Placeholder("Enter USD amount to convert").
				Value(&t.Amount),
		),
	)

	err := form.Run()
	if err != nil {
		return fmt.Errorf("form error: %w", err)
	}

	// Perform the conversion (mock implementation for now)
	//if err := t.performConversion(); err != nil {
	//	return err
	//}

	// Display the conversion result
	fmt.Printf("\n💱 %s %s is equal to %s %s\n", t.Amount, t.FromCurrency, t.Result, t.ToCurrency)

	return nil
}
