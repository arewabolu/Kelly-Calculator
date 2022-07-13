package kellycalc

import (
	"strconv"
	"strings"
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
	if strings.Contains(s, "%") {
		s = strings.TrimSuffix(s, "%")
	}
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
	Bal = Bal - 1.0
	return Bal, nil
}

func Fcalc(P, B float64) float64 {
	Q := 1.0 - P
	F := ((B * P) - Q) / B
	//F := P - (Q / B)
	return F
}

//246.25
func WithBal(Bal, F float64) float64 {
	use := F * Bal

	return use
}
