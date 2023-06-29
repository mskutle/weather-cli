# Weather CLI

Easily print the current weather + forecast for the rest of the day at any location.

## Prerequisites

Obtain a (free) API key from (weatherapi.com)[https://www.weatherapi.com/] and set the WEATHER_API_KEY env var on your system

## Installation (macOS)

1. Clone the repo and cd into the folder
2. `go build -o bin/weather`
3. Make it available from anywhere, `mv bin/weather /usr/local/bin`
4. Open up a terminal and execute it: `weather london`
