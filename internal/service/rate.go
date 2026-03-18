package service

import (
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/SANEKNAYMCHIK/mock-cbr/internal/models"
)

func GetRates(testID string) models.Rate {
	data := models.Rate{
		Valutes: []models.Valute{},
	}
	// Атрибуты, ни на что не влияющие (data.Date, data.Name)
	// Принимать под эти 2 поля -- параметры из эндпойнта
	// Не вижу смысла, так как просто перекопируются данные, которые не влияют на результат и логику
	// Поэтому для соответствия изначальной структуре xml, просто заполню эти поля заглушками
	data.Date = time.Now().Format("02.01.2006")
	data.Name = "Foreign Currency Market"
	seed := getSeed(testID)
	r := rand.New(rand.NewSource(seed))
	currentCurrencies := getRandomValutes(r, baseCurrencies)
	for _, item := range currentCurrencies {
		val := item
		testValue := randomize(r, val.Value)
		val.Value = pointToComma(testValue)
		if val.Nominal != 0 {
			val.VunitRate = pointToComma(testValue / float64(val.Nominal))
		}
		data.Valutes = append(data.Valutes, val)
	}
	return data
}

func getRandomValutes(r *rand.Rand, baseCurrencies []models.Valute) []models.Valute {
	res := []models.Valute{}
	for _, item := range baseCurrencies {
		if r.Float64() < 0.5 {
			res = append(res, item)
		}
	}
	if len(res) == 0 {
		res = append(res, baseCurrencies[0])
	}
	return res
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

// Чтобы данные не казались синтетическими,
// то используем базовый курс для каждой валюты (значение за какую-то дату) и возвращаем значение,
// не сильно отличающееся от базового

// randomize -- функция для рандомизации курса валюты,
// добавляющая к базовому курсу некий шум в пределах +-10% от базового курса
func randomize(r *rand.Rand, val string) float64 {
	baseVal, err := strconv.ParseFloat(strings.ReplaceAll(val, ",", "."), 64)
	if err != nil {
		baseVal = 0
	}
	return baseVal * (0.9 + r.Float64()*0.2)
}
