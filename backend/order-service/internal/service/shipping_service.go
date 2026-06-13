package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OSRMResponse struct {
	Routes []struct {
		Distance float64 `json:"distance"`
	} `json:"routes"`
}

func GetDistance(
	originLon string,
	originLat string,
	destLon string,
	destLat string,
) (float64, error) {

	url := fmt.Sprintf(
		"https://router.project-osrm.org/route/v1/driving/%s,%s;%s,%s?overview=false",
		originLon,
		originLat,
		destLon,
		destLat,
	)

	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var result OSRMResponse

	err = json.NewDecoder(
		resp.Body,
	).Decode(&result)

	if err != nil {
		return 0, err
	}

	return result.Routes[0].Distance / 1000, nil
}