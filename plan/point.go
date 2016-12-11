package plan

type Point struct {
	X int
	Y int
}

func (p *Point) IsEqual(point Point) bool {
	if p.X == point.X && p.Y == point.Y {
		return true
	}
	return false
}

func absInt(x int) (res int) {
	res = x
	if res < 0 {
		res = -res
	}
	return
}

func AddPoint(a *Point, b *Point) int {
	var res Point
	res.X = absInt(a.X - b.X)
	res.Y = absInt(a.Y - b.Y)
	return (res.X + res.Y)
}
