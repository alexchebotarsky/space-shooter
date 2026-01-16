package utils

import "testing"

func TestGetOriginPoint(t *testing.T) {
	type args struct {
		posX       float64
		posY       float64
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
			name: "Top-left corner, 0deg rotation",
			args: args{
				posX:       6,
				posY:       7,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0,
				originX:    0,
				originY:    0,
			},
			wantX: 6,
			wantY: 7,
		},
		{
			name: "Top-left corner, 90deg rotation",
			args: args{
				posX:       6,
				posY:       7,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0,
				originY:    0,
			},
			wantX: 11,
			wantY: 6,
		},
		{
			name: "Top-left corner, 180deg rotation",
			args: args{
				posX:       6,
				posY:       7,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.5,
				originX:    0,
				originY:    0,
			},
			wantX: 12,
			wantY: 11,
		},
		{
			name: "Top-left corner, 180deg rotation",
			args: args{
				posX:       6,
				posY:       7,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.5,
				originX:    0,
				originY:    0,
			},
			wantX: 12,
			wantY: 11,
		},
		{
			name: "Top-left corner, 270deg rotation",
			args: args{
				posX:       6,
				posY:       7,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   0.75,
				originX:    0,
				originY:    0,
			},
			wantX: 7,
			wantY: 12,
		},
		{
			name: "Top-left corner, 360deg rotation",
			args: args{
				posX:       6,
				posY:       7,
				rectWidth:  6,
				rectHeight: 4,
				rotation:   1,
				originX:    0,
				originY:    0,
			},
			wantX: 6,
			wantY: 7,
		},
		{
			name: "Top-right corner, 90deg rotation",
			args: args{
				posX:       3,
				posY:       4,
				rectWidth:  3,
				rectHeight: 4,
				rotation:   0.25,
				originX:    1,
				originY:    0,
			},
			wantX: 6.5,
			wantY: 7.5,
		},
		{
			name: "Bottom-right corner, 90deg rotation",
			args: args{
				posX:       3,
				posY:       4,
				rectWidth:  3,
				rectHeight: 4,
				rotation:   0.25,
				originX:    1,
				originY:    1,
			},
			wantX: 2.5,
			wantY: 7.5,
		},
		{
			name: "Bottom-left corner, 90deg rotation",
			args: args{
				posX:       3,
				posY:       4,
				rectWidth:  3,
				rectHeight: 4,
				rotation:   0.25,
				originX:    0,
				originY:    1,
			},
			wantX: 2.5,
			wantY: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := GetOriginPoint(tt.args.posX, tt.args.posY, tt.args.rectWidth, tt.args.rectHeight, tt.args.rotation, tt.args.originX, tt.args.originY)
			if gotX != tt.wantX {
				t.Errorf("GetOriginatedPoint() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("GetOriginatedPoint() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
