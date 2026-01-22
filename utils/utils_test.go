package utils

import (
	"testing"
)

func TestEuclideanMod(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive whole divisor",
			args: args{
				x: 10,
				y: 4,
			},
			want: 2,
		},
		{
			name: "Positive floating point divisor",
			args: args{
				x: 1.5,
				y: 1,
			},
			want: 0.5,
		},
		{
			name: "Negative whole divisor",
			args: args{
				x: 10,
				y: -4,
			},
			want: 2,
		},
		{
			name: "Negative floating point divisor",
			args: args{
				x: 1.5,
				y: -1,
			},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EuclideanMod(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("EuclideanMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlooredMod(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Positive whole divisor",
			args: args{
				x: 10,
				y: 4,
			},
			want: 2,
		},
		{
			name: "Positive floating point divisor",
			args: args{
				x: 1.5,
				y: 1,
			},
			want: 0.5,
		},
		{
			name: "Negative whole divisor",
			args: args{
				x: 10,
				y: -4,
			},
			want: -2,
		},
		{
			name: "Negative floating point divisor",
			args: args{
				x: 1.5,
				y: -1,
			},
			want: -0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlooredMod(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("FlooredMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundTo(t *testing.T) {
	type args struct {
		value     float64
		precision int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Round down",
			args: args{
				value:     3.721,
				precision: 2,
			},
			want: 3.72,
		},
		{
			name: "Round up",
			args: args{
				value:     3.729,
				precision: 2,
			},
			want: 3.73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundTo(tt.args.value, tt.args.precision); got != tt.want {
				t.Errorf("RoundTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectionMultipliers(t *testing.T) {
	tests := []struct {
		name      string
		direction float64
		wantMx    float64
		wantMy    float64
	}{
		{
			name:      "North",
			direction: 0,
			wantMx:    0,
			wantMy:    -1,
		},
		{
			name:      "North-East",
			direction: 0.125,
			wantMx:    0.5,
			wantMy:    -0.5,
		},
		{
			name:      "East",
			direction: 0.25,
			wantMx:    1,
			wantMy:    0,
		},
		{
			name:      "South-East",
			direction: 0.375,
			wantMx:    0.5,
			wantMy:    0.5,
		},
		{
			name:      "South",
			direction: 0.5,
			wantMx:    0,
			wantMy:    1,
		},
		{
			name:      "South-West",
			direction: 0.625,
			wantMx:    -0.5,
			wantMy:    0.5,
		},
		{
			name:      "West",
			direction: 0.75,
			wantMx:    -1,
			wantMy:    0,
		},
		{
			name:      "North-West",
			direction: 0.875,
			wantMx:    -0.5,
			wantMy:    -0.5,
		},
		{
			name:      "North (1)",
			direction: 1,
			wantMx:    0,
			wantMy:    -1,
		},
		{
			name:      "North (-1)",
			direction: -1,
			wantMx:    0,
			wantMy:    -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMx, gotMy := DirectionMultipliers(tt.direction)
			if gotMx != tt.wantMx {
				t.Errorf("DirectionMultipliers() gotMx = %v, want %v", gotMx, tt.wantMx)
			}
			if gotMy != tt.wantMy {
				t.Errorf("DirectionMultipliers() gotMy = %v, want %v", gotMy, tt.wantMy)
			}
		})
	}
}

func TestTransformedLocalPoint(t *testing.T) {
	type args struct {
		localX     float64
		localY     float64
		rectWidth  float64
		rectHeight float64
		rotation   float64
		originX    float64
		originY    float64
	}
	tests := []struct {
		name  string
		args  args
		wantX float64
		wantY float64
	}{
		{
			name: "Top-right, 0deg, origin center",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 6,
			wantY: 0,
		},
		{
			name: "Top-right, 90deg, origin center",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 5,
			wantY: 5,
		},
		{
			name: "Top-right, 180deg, origin center",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.5,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 0,
			wantY: 4,
		},
		{
			name: "Top-right, 270deg, origin center",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.75,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 1,
			wantY: -1,
		},
		{
			name: "Top-right, 360deg, origin center",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   1,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 6,
			wantY: 0,
		},
		{
			name: "Bottom-right, 90deg, origin center",
			args: args{
				localX:     6,
				localY:     4,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 1,
			wantY: 5,
		},
		{
			name: "Bottom-left, 90deg, origin center",
			args: args{
				localX:     0,
				localY:     4,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 1,
			wantY: -1,
		},
		{
			name: "Top-left, 90deg, origin center",
			args: args{
				localX:     0,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0.5,
				originY:    0.5,
			},
			wantX: 5,
			wantY: -1,
		},
		{
			name: "Top-right, 90deg, origin top-right",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    1,
				originY:    0,
			},
			wantX: 6,
			wantY: 0,
		},
		{
			name: "Top-right, 90deg, origin top-left",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0,
				originY:    0,
			},
			wantX: 0,
			wantY: 6,
		},
		{
			name: "Top-right, 90deg, origin bottom-left",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0,
				originY:    1,
			},
			wantX: 4,
			wantY: 10,
		},
		{
			name: "Top-right, 90deg, origin bottom-right",
			args: args{
				localX:     6,
				localY:     0,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    1,
				originY:    1,
			},
			wantX: 10,
			wantY: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := TransformedLocalPoint(tt.args.localX, tt.args.localY, tt.args.rectWidth, tt.args.rectHeight, tt.args.rotation, tt.args.originX, tt.args.originY)
			if gotX != tt.wantX {
				t.Errorf("LocalPointOnTransformedRect() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("LocalPointOnTransformedRect() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
