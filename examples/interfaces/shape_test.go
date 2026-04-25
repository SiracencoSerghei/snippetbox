package interfaces

import (
	"math"
	"testing"
	"fmt"
)

type shapeTestCase struct {
	name     string
	shape    Shape
	expected float64
}

func almostEqual(a, b float64) bool {
	const epsilon = 1e-6
	return math.Abs(a-b) < epsilon
}

func TestShapeArea(t *testing.T) {
	tests := []shapeTestCase{
		{
			name:     "circle radius 5",
			shape:    Circle{radius: 5},
			expected: math.Pi * 25,
		},
		{
			name:     "square length 6",
			shape:    Square{length: 6},
			expected: 36,
		},
		{
			name:     "triangle base 6 height 8",
			shape:    Triangle{base: 6, height: 8},
			expected: 24,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()

			// ВІЗУАЛІЗАЦІЯ (debug layer)
			t.Logf("index=%d type=%T value=%+v area=%.2f",
				i, tt.shape, tt.shape, got,
			)

			if !almostEqual(got, tt.expected) {
				t.Errorf("wrong area: got %v want %v", got, tt.expected)
			}
		})
	}
}

func TestShapeAreaWithInspector(t *testing.T) {
	inspector := NewInspector()

	shapes := []Shape{
		Circle{radius: 5},
		Square{length: 6},
		Triangle{base: 6, height: 8},
	}

	for i, shape := range shapes {
		t.Run(fmt.Sprintf("shape-%d", i), func(t *testing.T) {

			area := shape.Area()

			// 🔍 DEBUG INSPECTION
			inspector.Inspect(shape)

			t.Logf("area = %.2f", area)
		})
	}
}