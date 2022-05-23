package entity

type Satellite struct {
	Name  string
	Point *Point
}

func NewSatellite(name string, p *Point) *Satellite {
	return &Satellite{
		Name:  name,
		Point: p,
	}
}
