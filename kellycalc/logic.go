package kellycalc

import (
	"strconv"
)

//((BQ)-Q)/B=F

func BalCalc(s string) (float64, error) {
	Bal, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	return Bal, nil
}

func PCalc(s string) (float64, error) {
	P, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	P = P / 100
	return P, nil
}

func BCalc(s string) (float64, error) {
	Bal, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	return Bal, nil
}

func Fcalc(P, B float64) float64 {
	P = P / 100
	Q := 1 - P
	F := P - (Q / B)
	return F
}
