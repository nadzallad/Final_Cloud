package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type GeoResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func GetCoordinate(city string) (string, string, error) {

	api := fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		url.QueryEscape(city),
	)

	req, _ := http.NewRequest(
		"GET",
		api,
		nil,
	)

	req.Header.Set(
		"User-Agent",
		"order-service",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()

	var result []GeoResponse

	err = json.NewDecoder(
		resp.Body,
	).Decode(&result)

	if err != nil {
		return "", "", err
	}

	if len(result) == 0 {
		return "", "", fmt.Errorf("city not found")
	}

	return result[0].Lat,
		result[0].Lon,
		nil
}
