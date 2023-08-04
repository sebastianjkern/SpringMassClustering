package main

import "SpringMassClustering/proto"

type Spring struct {
	RelaxedLength  float32
	SpringConstant float32
	Mass1          *Mass
	Mass2          *Mass
}

func NewSpringFromPB(spring *proto.Spring) *Spring {
	return &Spring{
		RelaxedLength:  spring.RelaxedLength,
		SpringConstant: spring.SpringConstant,
	}
}

func (s *Spring) Update() {
	direction := s.Mass2.Position.Minus(s.Mass1.Position)
	normalized := direction.DivideBy(direction.Distance())

	resultingForce := normalized.MultiplyScalar(direction.Distance() - s.RelaxedLength).MultiplyScalar(s.SpringConstant)

	minus := s.Mass2.Force.Minus(resultingForce)
	plus := s.Mass1.Force.Plus(resultingForce)

	if s.Mass1.Fixed && s.Mass2.Fixed {
		return
	} else if s.Mass1.Fixed {
		s.Mass2.Force = minus
	} else if s.Mass2.Fixed {
		s.Mass1.Force = plus
	}

	s.Mass2.Force = minus.MultiplyScalar(0.5)
	s.Mass1.Force = plus.MultiplyScalar(0.5)
}
