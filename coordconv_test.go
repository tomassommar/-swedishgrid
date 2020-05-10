package swedishgrid

import "testing"

func TestRT90toWGS84(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{"Test reference point 1", args{7453389.762, 1727060.905}, 67.09068132460693, 21.034750437140964},
		{"Test reference point 2", args{7047738.415, 1522128.637}, 63.53743381535966, 16.249887395670605},
		{"Test reference point 3", args{6671665.273, 1441843.186}, 60.15941178140019, 14.757824208548275},
		{"Test reference point 4", args{6249111.351, 1380573.079}, 56.354777568077445, 13.873265005640663},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RT90toWGS84(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("RT90toWGS84() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RT90toWGS84() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
