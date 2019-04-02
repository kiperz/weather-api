package client

import (
	weatherapi "github.com/briandowns/openweathermap"
)

type Client struct {
	key string
}

func New(key string) *Client {
	return &Client{
		key: key,
	}
}

type WeatherData struct {
	Type string
	Temperature float64
	MinimumTemperature float64
	MaximumTemperature float64
}

func (client *Client) GetWeatherData(locationName string) (*WeatherData, error) {
	var (
		res *weatherapi.CurrentWeatherData
		err error
	)
	if res, err = weatherapi.NewCurrent("C", "EN", client.key); err != nil {
		return nil, err
	}

	if err = res.CurrentByName(locationName); err != nil {
		return nil, err
	}

	return &WeatherData{
		Type: res.Weather[0].Description,
		Temperature: res.Main.Temp,
		MinimumTemperature: res.Main.TempMin,
		MaximumTemperature: res.Main.TempMax,
	}, nil
}