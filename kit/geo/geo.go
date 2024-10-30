package geo

import "math"

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Earth radius in kilometer
	const R = 6371

	// convert latitude and longitude from degrees to radians
	lat1Rad := radians(lat1)
	lon1Rad := radians(lon1)
	lat2Rad := radians(lat2)
	lon2Rad := radians(lon2)

	// difference in coordinates
	dLat := lat1Rad - lat2Rad
	dLon := lon1Rad - lon2Rad

	// haversine formula
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// distance in kilometers
	distance := R * c

	return distance
}

func SphericalLawofCosines(lat1, lon1, lat2, lon2 float64) float64 {
	// Earth radius in kilometer
	const R = 6371

	// convert latitude and longitude from degrees to radians
	lat1Rad := radians(lat1)
	lon1Rad := radians(lon1)
	lat2Rad := radians(lat2)
	lon2Rad := radians(lon2)

	// Calculate the cosine of the central angle
	cosC := math.Sin(lat1Rad)*math.Sin(lat2Rad) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Cos(lon2Rad-lon1Rad)

	// Avoid domain error due to floating point arithmetic
	if cosC > 1 {
		cosC = 1
	} else if cosC < -1 {
		cosC = -1
	}

	// Calculate the central angle in radians
	c := math.Acos(cosC)

	// distance in kilometers
	distance := R * c

	return distance
}

func radians(deg float64) float64 {
	return deg * math.Pi / 180
}
