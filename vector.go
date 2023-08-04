package main

import "math"

type Vector []float32

func (vector *Vector) Distance() float32 {
	sum := 0.0
	for i := 0; i < len(*vector); i++ {
		sum += math.Pow(float64((*vector)[i]), 2)
	}
	return float32(math.Sqrt(sum))
}

func (vector Vector) Minus(other Vector) Vector {
	result := Vector{}

	for i := 0; i < len(vector); i++ {
		result = append(result, (vector)[i]-(other)[i])
	}

	return result
}

func (vector *Vector) Plus(other Vector) Vector {
	result := Vector{}

	for i := 0; i < len(*vector); i++ {
		result = append(result, (*vector)[i]+(other)[i])
	}

	return result
}

func (vector Vector) DivideBy(value float32) Vector {
	result := Vector{}

	for i := 0; i < len(vector); i++ {
		result = append(result, (vector)[i]/value)
	}

	return result
}

func (vector Vector) MultiplyScalar(value float32) Vector {
	result := Vector{}

	for i := 0; i < len(vector); i++ {
		result = append(result, (vector)[i]*value)
	}

	return result
}
