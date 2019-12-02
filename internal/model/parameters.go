package model

const defaultEta = 0.5
const defaultMu = 0.5

// Parameters for neural network training.
type Parameters struct {
	Eta float64
	Mu  float64
}

// SetEta to the given value.
func (p *Parameters) SetEta(v float64) {
	p.Eta = v
}

// SetMu to the given value.
func (p *Parameters) SetMu(v float64) {
	p.Mu = v
}