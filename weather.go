package main

type Weather struct {
	Location Location `json:"location"`
	Current  struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast Forecast `json:"forecast"`
}

type Location struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type Forecast struct {
	Day []struct {
		Hour []struct {
			TimeEpoch int64   `json:"time_epoch"`
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
			ChanceOfRain float64 `json:"chance_of_rain"`
		} `json:"hour"`
	} `json:"forecastday"`
}
