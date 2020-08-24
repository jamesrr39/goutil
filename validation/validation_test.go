package validation

import "testing"

func TestAssertInRangeFloat64(t *testing.T) {
	type args struct {
		rangeValidation RangeValidation
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"in range",
			args{
				rangeValidation: RangeValidation{
					Value:      0.5,
					LowerBound: 0,
					UpperBound: 1,
				},
			},
			false,
		},
		{
			"too high",
			args{
				rangeValidation: RangeValidation{
					Value:      1.1,
					LowerBound: 0,
					UpperBound: 1,
				},
			},
			true,
		},
		{
			"too low",
			args{
				rangeValidation: RangeValidation{
					Value:      -0.1,
					LowerBound: 0,
					UpperBound: 1,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AssertInRangeFloat64(tt.args.rangeValidation); (err != nil) != tt.wantErr {
				t.Errorf("AssertInRangeFloat64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
