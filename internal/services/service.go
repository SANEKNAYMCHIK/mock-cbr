package services

import "github.com/SANEKNAYMCHIK/mock-cbr/internal/models"

func GetRates(testId string, statusCode int) models.Rate {
	data := models.Rate{
		Valutes: []models.Valute{
			{
				ValuteID: "010101",
				NumCode:  "001",
				CharCode: "RUB",
				Nominal:  100,
				Name:     "Российский рубль",
				Value:    1.00,
			},
			{
				ValuteID: "ASB215",
				NumCode:  "231",
				CharCode: "BYR",
				Nominal:  1000,
				Name:     "Белорусский рубль",
				Value:    154.00,
			},
			{
				ValuteID: "D94GHN1",
				NumCode:  "015",
				CharCode: "KRW",
				Nominal:  1,
				Name:     "Вон",
				Value:    2.2554,
			},
		},
	}
	return data
}
