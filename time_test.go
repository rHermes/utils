package utils

import (
	"testing"
	"time"
)

func TestUnixDuration(t *testing.T) {
	type TestCase struct {
		Name     string
		Input    time.Duration
		Expected time.Time
	}

	zt := time.Unix(0, 0)
	// Good cases
	cases := []TestCase{
		TestCase{
			Name:     "Simple seconds",
			Input:    time.Second * 10,
			Expected: zt.Add(time.Second * 10),
		},
		TestCase{
			Name:     "Microseconds",
			Input:    time.Microsecond * 101,
			Expected: zt.Add(time.Microsecond * 101),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			if cd := UnixDuration(tc.Input); !cd.Equal(tc.Expected) {
				t.Errorf("Expected %s, but got %s\n", tc.Expected.String(), cd.String())
			}
		})
	}

}
