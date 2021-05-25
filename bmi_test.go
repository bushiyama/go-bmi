package bmi

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMain(t *testing.T) {
	cases := []struct {
		name  string
		param []string
		want  int
	}{
		{
			name:  "正常1",
			param: []string{"90", "180", "22"},
			want:  19,
		},
		{
			name:  "正常2",
			param: []string{"94", "175", "22"},
			want:  27,
		},
		{
			name:  "under 1",
			param: []string{"66", "175", "22"},
			want:  0,
		},
		{
			name:  "eq",
			param: []string{"67", "175", "22"},
			want:  0,
		},
		{
			name:  "over",
			param: []string{"68", "175", "22"},
			want:  1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ret := bmi(tc.param)
			assert.Equal(t, ret, tc.want)
		})
	}
}
