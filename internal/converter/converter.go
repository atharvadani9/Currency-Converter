package converter

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/huh"
	"io"
	"net/http"
	"os"
	"strconv"
)

type ExchangeRatesResponse struct {
	Disclaimer  string             `json:"disclaimer"`
	License     string             `json:"license"`
	Timestamp   int64              `json:"timestamp"`
	Base        string             `json:"base"`
	Rates       map[string]float64 `json:"rates"`
	Error       bool               `json:"error"`
	Description string             `json:"description"`
}

type Tui struct {
	FromCurrency string
	ToCurrency   string
	Amount       string
	Result       string
}

func (t *Tui) CurrencyConverter() error {
	t.FromCurrency = "USD"

	currencies := []huh.Option[string]{
		huh.NewOption("Euro", "EUR"),
		huh.NewOption("British Pound", "GBP"),
		huh.NewOption("Indian Rupee", "INR"),
		huh.NewOption("New Zealand Dollar", "NZD"),
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

	if err := t.performConversion(); err != nil {
		return err
	}

	fmt.Printf("\n💱 %s %s is equal to %s %s\n", t.Amount, t.FromCurrency, t.Result, t.ToCurrency)

	return nil
}

func (t *Tui) performConversion() error {
	amount, err := strconv.ParseFloat(t.Amount, 64)
	if err != nil {
		return fmt.Errorf("invalid amount: %w", err)
	}

	appID := os.Getenv("APP_ID")
	if appID == "" {
		return fmt.Errorf("APP_ID environment variable is not set. Please set your OpenExchangeRates API key")
	}

	// Make API call to OpenExchangeRates
	apiURL := fmt.Sprintf("https://openexchangerates.org/api/historical/2013-02-16.json?app_id=%s", appID)
	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to make API request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close response body: %w", err)
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var exchangeRates ExchangeRatesResponse
	if err := json.Unmarshal(body, &exchangeRates); err != nil {
		return fmt.Errorf("failed to parse JSON response: %w", err)
	}

	if exchangeRates.Error {
		return fmt.Errorf("API request failed: %s", exchangeRates.Description)
	}
	rate, exists := exchangeRates.Rates[t.ToCurrency]
	if !exists {
		return fmt.Errorf("exchange rate not found for currency: %s", t.ToCurrency)
	}

	convertedAmount := amount * rate
	t.Result = fmt.Sprintf("%.2f", convertedAmount)

	return nil
}
