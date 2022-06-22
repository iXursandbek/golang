package lecture01

import "testing"

func TestCalcSquare(t *testing.T) {
	got := CalcSquare(10.0, SidesTriangle)
	want := 43.30127018922193

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

	got1 := CalcSquare(10.0, SidesSquare)
	want1 := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got1, want1)
	}

	got2 := CalcSquare(10.0, SidesCircle)
	want2 := 314.1592653589793

	if got != want {
		t.Errorf("got %.2f want %.2f", got2, want2)
	}

	got3 := CalcSquare(10.0, 5)
	want3 := 0.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got3, want3)
	}

}
