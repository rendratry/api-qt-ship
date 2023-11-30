package helper

import (
	"qt-api/model/domain"
	"strconv"
	"strings"
)

var data = [][]string{
	{"1", "Kapan Container 1", "20 knots", "Surabaya to Jakarta", "-39384938", "38748374"},
	{"2", "Kapan Container 2", "10 knots", "Surabaya to Pontianak", "-39384938", "38748374"},
	{"3", "Kapan Container 3", "21 knots", "Surabaya to Batam", "-39384938", "38748374"},
	{"4", "Kapan Container 3", "11 knots", "Surabaya to Jakarta", "-39384938", "38748374"},
}

func FilterData(id, isAll string) []domain.Ship {
	var result []domain.Ship

	if isAll == "true" {
		for _, shipData := range data {
			shipId, err := strconv.Atoi(shipData[0])
			PanicIfError(err)
			result = append(result, domain.Ship{
				ID:          shipId,
				Name:        shipData[1],
				Speed:       shipData[2],
				Destination: shipData[3],
				CurrentLocation: domain.CurrentLocation{
					Latitude:  shipData[4],
					Longitude: shipData[5],
				},
			})
		}
		return result
	}

	for _, shipData := range data {
		if id == "" || strings.EqualFold(shipData[0], id) {
			shipId, err := strconv.Atoi(shipData[0])
			PanicIfError(err)
			result = append(result, domain.Ship{
				ID:          shipId,
				Name:        shipData[1],
				Speed:       shipData[2],
				Destination: shipData[3],
				CurrentLocation: domain.CurrentLocation{
					Latitude:  shipData[4],
					Longitude: shipData[5],
				},
			})
		}
	}

	return result
}
