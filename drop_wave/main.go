package main

import (
	"fmt"
	m "math"
	"math/rand"

	"github.com/MaxHalford/gago"
)

// A Vector contains float64s.
type Vector []float64

// Evaluate a Vector with the Drop-Wave function which takes two variables as
// input and reaches a minimum of -1 in X = (0, 0).
func (X Vector) Evaluate() (float64, error) {
	var (
		numerator   = 1 + m.Cos(12*m.Sqrt(m.Pow(X[0], 2)+m.Pow(X[1], 2)))
		denominator = 0.5*(m.Pow(X[0], 2)+m.Pow(X[1], 2)) + 2
	)
	return -numerator / denominator, nil
}

// Mutate a Vector by applying by resampling each element from a normal
// distribution with probability 0.8.
func (X Vector) Mutate(rng *rand.Rand) {
	gago.MutNormalFloat64(X, 0.8, rng)
}

// Crossover a Vector with another Vector by applying uniform crossover.
func (X Vector) Crossover(Y gago.Genome, rng *rand.Rand) {
	gago.CrossUniformFloat64(X, Y.(Vector), rng)
}

// Clone a Vector.
func (X Vector) Clone() gago.Genome {
	var XX = make(Vector, len(X))
	copy(XX, X)
	return XX
}

// MakeVector returns a random vector by generating 2 values uniformally
// distributed between -10 and 10.
func MakeVector(rng *rand.Rand) gago.Genome {
	return Vector(gago.InitUnifFloat64(2, -10, 10, rng))
}

func main() {
	var ga = gago.Generational(MakeVector)
	ga.NPops = 1
	ga.Initialize()

	fmt.Printf("Best fitness at generation 0: %f\n", ga.HallOfFame[0].Fitness)
	for i := 0; i < 30; i++ {
		ga.Evolve()
		fmt.Printf("Best fitness at generation %d: %f\n", i, ga.HallOfFame[0].Fitness)
	}
}
