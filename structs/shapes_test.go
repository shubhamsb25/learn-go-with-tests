package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := Perimeter(rectangle)
	want := 60.0

	compareValues(got, want, t)
}

func compareValues(got float64, want float64, t *testing.T) {
	if got != want {
		t.Errorf(" got %g want %g", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Length: 10.0, Width: 20.0}, want: 200.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10.0, Height: 10.0}, want: 50.0},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.want {
				t.Errorf("%#v got %g want %g", areaTest.shape, got, areaTest.want)
			}
		})
	}
}
