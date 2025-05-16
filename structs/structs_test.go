package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Length: 10.0, Width: 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Length: 10, Width: 10}, hasArea: 100.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.hasArea)
	}
}
