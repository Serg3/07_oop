package hw

import "testing"

func TestGeom_Distance(t *testing.T) {
	tests := []struct {
		name         string
		geom         Geom
		wantDistance float64
	}{
		{
			name:         "#1",
			geom:         Geom{x1: 1, y1: 1, x2: 4, y2: 5},
			wantDistance: 5,
		},
		{
			name:         "#2",
			geom:         Geom{x1: 0, y1: 0, x2: 0, y2: 0},
			wantDistance: 0,
		},
		{
			name:         "#3",
			geom:         Geom{x1: 1, y1: 1, x2: -4, y2: 5},
			wantDistance: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := tt.geom.Distance(); gotDistance != tt.wantDistance {
				t.Errorf("Geom.Distance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
