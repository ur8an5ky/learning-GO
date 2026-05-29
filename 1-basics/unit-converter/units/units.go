package units

import (
	"fmt"
	"strconv"
)

type Category int

const (
	CategoryUnknown Category = iota - 1
	CategoryTemperature
	CategoryLength
	CategoryWeight
)

type Units struct {
	temperature map[string]float64

	length map[string]float64

	weight map[string]float64
}

func NewUnits() Units {
	return Units{
		temperature: map[string]float64{
			"c": 1.0,
			"f": 32.0,
			"k": 273.15,
		},
		length: map[string]float64{
			"m":  1.0,
			"km": 1000.0,
			"cm": 0.01,
			"mi": 1609.344,
			"ft": 0.3048,
		},
		weight: map[string]float64{
			"g":  1.0,
			"kg": 1000.0,
			"t":  1000000.0,
			"mg": 0.001,
			"lb": 453.59237,
			"oz": 28.349523125,
		},
	}
}

func bothInMap(m map[string]float64, u1, u2 string) bool {
	_, ok1 := m[u1]
	_, ok2 := m[u2]
	return ok1 && ok2
}

func TypeOfCategory(unit1, unit2 string) Category {
	u := NewUnits()
	switch {
	case bothInMap(u.temperature, unit1, unit2):
		return CategoryTemperature
	case bothInMap(u.length, unit1, unit2):
		return CategoryLength
	case bothInMap(u.weight, unit1, unit2):
		return CategoryWeight
	default:
		return CategoryUnknown
	}
}

func (u Units) ConvertTemperature(command []string) (float64, error) {
	value, err := strconv.ParseFloat(command[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", command[0], err)
	}

	celsius := toCelsius(value, command[1])
	return fromCelsius(celsius, command[3]), nil
}

func toCelsius(value float64, unit string) float64 {
	switch unit {
	case "c":
		return value
	case "k":
		return value - 273.15
	case "f":
		return (value - 32) * 5 / 9
	default:
		return value
	}
}

func fromCelsius(value float64, unit string) float64 {
	switch unit {
	case "c":
		return value
	case "k":
		return value + 273.15
	case "f":
		return value*9/5 + 32
	default:
		return value
	}
}

func (u Units) ConvertLength(command []string) (float64, error) {
	value, err := strconv.ParseFloat(command[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", command[0], err)
	}
	unit1, unit2 := u.length[command[1]], u.length[command[3]]

	return value / unit2 * unit1, nil
}

func (u Units) ConvertWeight(command []string) (float64, error) {
	value, err := strconv.ParseFloat(command[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", command[0], err)
	}
	unit1, unit2 := u.weight[command[1]], u.weight[command[3]]

	return value / unit2 * unit1, nil
}
