# Weather CLI

Easily print the current weather + forecast for the rest of the day at any location.

## Prerequisites

Obtain a (free) API key from (weatherapi.com)[https://www.weatherapi.com/] and set the WEATHER_API_KEY env var on your system

## Installation (on macOS. Should be similar on other OSes)

1. Clone the repo and cd into the folder
2. `go build -o bin/weather`
3. Make it available from anywhere, `mv bin/weather /usr/local/bin`
4. Open up a terminal and execute it: `weather london`

Output ->

London, United Kingdom: 17C, Partly cloudy
01:00 - 18C, 86%, Light rain shower
02:00 - 18C, 64%, Light rain shower
03:00 - 18C, 77%, Light rain shower
04:00 - 17C, 81%, Moderate rain
05:00 - 17C, 64%, Light rain shower
06:00 - 15C, 87%, Light rain shower
07:00 - 14C, 87%, Moderate rain
08:00 - 14C, 84%, Moderate rain
09:00 - 14C, 83%, Light rain
10:00 - 15C, 66%, Moderate rain
11:00 - 16C, 79%, Light rain
12:00 - 18C, 76%, Patchy light drizzle
13:00 - 20C, 68%, Moderate rain
14:00 - 21C, 0%, Cloudy
15:00 - 20C, 0%, Overcast
16:00 - 20C, 0%, Partly cloudy
17:00 - 20C, 0%, Partly cloudy
18:00 - 21C, 0%, Partly cloudy
19:00 - 20C, 0%, Overcast
20:00 - 20C, 0%, Partly cloudy
21:00 - 18C, 0%, Partly cloudy
22:00 - 17C, 0%, Partly cloudy
23:00 - 16C, 0%, Partly cloudy
00:00 - 15C, 0%, Clear
