package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type WeatherData struct {
	City    string  `json:"name"`
	Country string  `json:"sys.country"`
	Temp    float64 `json:"main.temp"`
}

func fetchWeatherData(city string, wg *sync.WaitGroup) {
	defer wg.Done()

	apiKey := "YOUR_API_KEY"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather data for %s: %v\n", city, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: %v\n", city, err)
		return
	}

	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("Error decoding JSON for %s: %v\n", city, err)
		return
	}

	fmt.Printf("Weather in %s, %s: %.2f°C\n", weather.City, weather.Country, weather.Temp-273.15)
}

func main() {
	cities := []string{"London", "Paris", "New York", "Tokyo", "Moscow"}

	var wg sync.WaitGroup
	for _, city := range cities {
		wg.Add(1)
		go fetchWeatherData(city, &wg)
	}
	wg.Wait()
}
