package geo_test

import (
	"test_cases/kit/geo"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TEST BENCHMARK

func TestDistance(t *testing.T) {
	type args struct {
		lat1 float64
		lon1 float64
		lat2 float64
		lon2 float64
	}
	testCases := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "calculate distance between monas and stasiun kota",
			args: args{
				lat1: 5.570877512722721,
				lon1: 95.36953887460648,
				lat2: 5.579063119868789,
				lon2: 95.36759873387135,
			},
			want: 0.9,
		},
		{
			name: "calculate distance stasiun kota and somewhere else",
			args: args{
				lat1: 5.570877512722721,
				lon1: 95.36953887460648,
				lat2: 5.580627769907465,
				lon2: 95.358051482646,
			},
			want: 1.5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// if got := geo.Distance(tc.args.lat1, tc.args.lon1, tc.args.lat2, tc.args.lon2); got != tc.want {
			// 	t.Errorf("Distance() = %v, want %v", got, tc.want)
			// }

			got := geo.Haversine(tc.args.lat1, tc.args.lon1, tc.args.lat2, tc.args.lon2)
			assert.GreaterOrEqual(t, got, tc.want, "the distance should be greater or equal the expected number")
		})
	}
}

// menjalankan semua benchmark
// go test ./kit/geo/geo_test.go -bench=. -benchmem -run=none

// menjalankan salahsatu benchmark
// go test ./kit/geo/geo_test.go -bench=BenchmarkHaversine -benchmem -run=none

func BenchmarkHaversine(b *testing.B) {
	for i := 0; i < 1000; i++ {
		geo.Haversine(5.570877512722721, 95.36953887460648, 5.580627769907465, 95.358051482646)
	}
}

func BenchmarkSphericalLawofCosines(b *testing.B) {
	for i := 0; i < 1000; i++ {
		geo.SphericalLawofCosines(5.570877512722721, 95.36953887460648, 5.580627769907465, 95.358051482646)
	}
}
