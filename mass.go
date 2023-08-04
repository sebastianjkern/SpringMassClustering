package main

import "SpringMassClustering/proto"

type Mass struct {
	Id       int32
	Fixed    bool
	Position Vector
	Weight   float32
	Speed    Vector
	Force    Vector
	Drag     float32
}

func NewMassFromPB(mass *proto.Mass) *Mass {
	return &Mass{
		Id:       mass.Id,
		Fixed:    mass.Fixed,
		Position: []float32{mass.PosX, mass.PosY},
		Weight:   mass.Weight,
		Speed:    []float32{0, 0},
		Force:    []float32{0, 0},
		Drag:     mass.Drag,
	}
}

func (m *Mass) Update(dt float32) {
	if !m.Fixed {
		m.Force = m.Force.Minus(m.Speed.MultiplyScalar(m.Drag))
		m.Speed = m.Speed.Plus(m.Force.DivideBy(m.Weight).MultiplyScalar(dt))
		m.Position = m.Position.Plus(m.Speed.MultiplyScalar(dt))
	}
}
