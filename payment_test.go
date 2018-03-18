package financial

import (
	"testing"
)

const (
	rate     = .056
	nper     = 36.
	pv       = 35698.
	fv       = 10000.
	expected = 823.80
)

func Test(t *testing.T) {

	payment := PMT(rate/12, nper, pv, fv)

	if payment != expected {
		t.Errorf("pmt %v", payment)
	}
}
