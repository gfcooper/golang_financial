package financial

import (
	"testing"
)

func Test01(t *testing.T) {

	payment := PMT(0.056/12, 36., 15000., 5000., 0)
	expects := 325.74

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}

func Test02(t *testing.T) {

	payment := PMT(0/12, 24., 15000., 5000., 0)
	expects := 416.67

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}

func Test03(t *testing.T) {

	payment := PMT(0.056/12, 36., 15000., 0., 0)
	expects := 453.62

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}

func Test04(t *testing.T) {

	payment := PMT(0/12, 24., 15000., 0., 0)
	expects := 625.

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}

func Test05(t *testing.T) {

	payment := PMT(0.056/12, 36., 15000., 5000., 1)
	expects := 324.23

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}

func Test06(t *testing.T) {

	payment := PMT(0/12, 24., 15000., 5000., 1)
	expects := 416.67

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}

func Test07(t *testing.T) {

	payment := PMT(0.056/12, 36., 15000., 0., 1)
	expects := 451.51

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}

func Test08(t *testing.T) {

	payment := PMT(0/12, 24., 15000., 0., 1)
	expects := 625.

	if payment != expects {
		t.Errorf("pmt %v", payment)
	}
}
