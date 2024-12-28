package data

type location string

func DistanceTo(destination location) distance {
	//TODO calculations...
	return 10
}

type distance float64 //miles
type distanceKm float64

// Method
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (miles distanceKm) ToMiles() distance {
	return distance(miles / 1.60934)
}

func test() {
	d := distance(34.5)
	km := d.ToKm()
	km.ToMiles()
	print(d)
}
