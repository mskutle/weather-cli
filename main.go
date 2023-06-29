package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")

	if len(apiKey) == 0 {
		panic("missing api key")
	}

	q := "Fiskum"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	resp, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1", apiKey, q))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Day[0].Hour

	fmt.Printf(
		"%s, %s: %0.fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		// if date.Before(time.Now()) {
		// 	continue
		// }

		msg := fmt.Sprintf(
			"%s - %0.fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(msg)
			continue
		}

		color.Red(msg)
	}
}
