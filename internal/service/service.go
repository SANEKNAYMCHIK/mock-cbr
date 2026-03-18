package service

import (
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"

	"github.com/SANEKNAYMCHIK/mock-cbr/internal/models"
)

func GetRates(testID string, statusCode int) models.Rate {
	data := models.Rate{
		Valutes: []models.Valute{},
	}
	// Атрибуты, ни на что не влияющие (data.Date, data.Name)
	// Принимать под эти 2 поля -- параметры из эндпойнта
	// Не вижу смысла, так как просто перекопируются данные, которые не влияют на результат и логику
	// Поэтому для соответствия изначальной структуре xml, просто заполню эти поля заглушками
	data.Date = "18.03.2026"
	data.Name = "Foreign Currency Market"
	seed := getSeed(testID)
	r := rand.New(rand.NewSource(seed))
	for _, item := range baseCurrencies {
		val := item
		testValue := randomize(r)
		val.Value = pointToComma(testValue)
		if val.Nominal != 0 {
			val.VunitRate = pointToComma(testValue / float64(val.Nominal))
		}
		data.Valutes = append(data.Valutes, val)
	}
	return data
}

func getSeed(testID string) int64 {
	h := fnv.New64a()
	h.Write([]byte(testID))
	return int64(h.Sum64())
}

func pointToComma(val float64) string {
	floatString := strconv.FormatFloat(val, 'f', 4, 64)
	return strings.ReplaceAll(floatString, ".", ",")
}

func randomize(r *rand.Rand) float64 {
	return r.Float64() * 100
}
