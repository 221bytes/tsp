package plan

type Location struct {
	lat  float32
	long float32
}

func (p *Location) IsEqual(Location Location) bool {
	if p.lat == Location.lat && p.long == Location.long {
		return true
	}
	return false
}

func absFloat(lat float32) (res float32) {
	res = lat
	if res < 0 {
		res = -res
	}
	return
}

func AddLocation(a *Location, b *Location) float32 {
	var res Location
	res.lat = absFloat(a.lat - b.lat)
	res.long = absFloat(a.long - b.long)
	return (res.lat + res.long)
}
