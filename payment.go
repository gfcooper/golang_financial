package financial

import (
	"math"
)

// PMT calculates a payment close enough for what we need.
func PMT(rate, periods, principal, balloon float64) (payment float64) {

	pPayment := principalPayment(rate, periods, principal-balloon)

	bPayment := balloonPayment(rate, balloon)

	payment = (pPayment + bPayment) * 100

	return math.Round(payment) / 100
}

func principalPayment(rate, periods, pminb float64) (payment float64) {

	if rate == 0 {
		payment = (pminb / periods)
	} else {
		payment = (pminb * (rate * math.Pow(1+rate, periods))) / (math.Pow(1+rate, periods) - 1)
	}

	return
}

func balloonPayment(rate, balloon float64) (payment float64) {

	if rate == 0 || balloon == 0 {
		payment = 0
	} else {
		payment = balloon * rate
	}

	return
}
