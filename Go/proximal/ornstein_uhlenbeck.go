package main

import (
	"image/color"
	"sort"
	"fmt"

	"math"
	"math/rand"
	//"time"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
//	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var rnd *rand.Rand

func sortVectorInOrder (v mat.Vector, order []int) mat.Vector {
	n := v.Len ()
	if n != len (order) {
		panic ("Doesn't match the length of arrays")
	}

	data := make ([]float64, n)
	for i := range (data) {
		data[i] = v.AtVec (order[i])
	}

	return mat.NewVecDense (n, data)
}

func argsortVector (v mat.Vector) []int {
	n := v.Len()
	data := make ([]float64, n)
	order := make ([]int, n)
	for i := range(data) {
		data[i] = v.AtVec (i)
	}
	floats.Argsort (data, order)
	return order
}

func checkSize (sigma, psi mat.Vector, C mat.Matrix) bool {
	c0 := sigma.Len()
	c1 := psi.Len()
	c2, r2 := C.Dims()

	return c0 != c1 || c1 != c2 || c2 != r2
}

func euclideanDistance (a, b mat.Vector) float64 {
	var v mat.VecDense
	v.SubVec (a, b)

	return math.Sqrt (mat.Dot(&v, &v))
}


func proxRecur (sigma, psi mat.Vector, C *mat.Dense, beta, h float64) *mat.VecDense {
	epsilon := 0.05	// regularization parameter
	delta := 0.001	// tolerance
	L := 100		// max iteration

	if checkSize(sigma, psi, C) {
		panic ("Doesn't match the length of arrays")
	}

	N, _ := C.Dims() // # of samples

	// initialize
	var gamma, xi, tmp mat.Dense
	var vtmp mat.VecDense
	z := make ([]mat.Vector, L)
	y := make ([]mat.Vector, L)

	gamma.Apply (
		func(i, j int, v float64) float64 { return math.Exp (-v/(2*epsilon)) },
		C)
	xi.Apply (
		func(i, j int, v float64) float64 { return math.Exp (-beta*v-1) },
		psi)

	z0 := mat.NewVecDense (N, nil)
	for col:=0; col<N; col++ {
		z0.SetVec (col, rnd.Float64())
	}
	z[0] = z0
	tmp.Mul (&gamma, z[0])
	tmp.DivElem (sigma, &tmp)
	y[0] = mat.VecDenseCopyOf (tmp.ColView(0))

	// main loop
	for l:=1; l<L; l++ {
		// calculate z
		tmp.Reset()
		tmp.Mul (gamma.T(), y[l-1])
		tmp.DivElem (&xi, &tmp)
		tmp.Apply (
			func(i,j int, v float64) float64 { return math.Pow (v, 1.0/(1.0+beta*epsilon/h)) }, 
			&tmp)
		z[l] = mat.VecDenseCopyOf (tmp.ColView(0))

		// calculate y
		tmp.Reset()
		tmp.Mul (&gamma, z[l])
		tmp.DivElem (sigma, &tmp)
		y[l] = mat.VecDenseCopyOf (tmp.ColView(0))

		// check tolerance
		if euclideanDistance(y[l-1], y[l]) < delta &&
			euclideanDistance(z[l-1], z[l]) < delta {

			var newSigma mat.VecDense
			vtmp.MulVec (gamma.T(), y[l])
			newSigma.MulElemVec (z[l], &vtmp)
			return &newSigma
		}
	}

	var newSigma mat.VecDense
	vtmp.MulVec (gamma.T(), y[L-1])
	newSigma.MulElemVec (z[L-1], &vtmp)
	return &newSigma
}

func eulerMaruyama (initVals, initDistbs mat.Vector, a, beta, h float64) ([]mat.Vector, []mat.Vector) {
	nSamples := initVals.Len()
	maxIter := 4000
	// maxIter := 10
	stateVals := make ([]mat.Vector, maxIter+1)
	distbs := make ([]mat.Vector, maxIter+1)

	// sort
	order := argsortVector (initVals)

	stateVals[0] = sortVectorInOrder (initVals, order)
	distbs[0] = sortVectorInOrder (initDistbs, order)

	for iter:=0; iter<maxIter; iter++ {
		// update state
		var x, dx, tmp1 mat.VecDense
		tmp1.ScaleVec (-a*h, stateVals[iter])
		rnddata := make ([]float64, nSamples)
		for col:=0; col<nSamples; col++ {
			rnddata[col] = math.Sqrt (2.0*h/beta) * rnd.NormFloat64()
		}
		tmp2 := mat.NewVecDense (nSamples, rnddata)

		dx.AddVec (&tmp1, tmp2)
		x.AddVec (stateVals[iter], &dx)

		order := argsortVector (&x)
		stateVals[iter+1] = sortVectorInOrder (&x, order)

		// proximal recursion
		rawPsi := make ([]float64, nSamples)
		rawC := make ([]float64, nSamples*nSamples)
		for col:=0; col<nSamples; col++ {
			xe := stateVals[iter+1].AtVec (col)
			rawPsi[col] = 0.5*a*xe*xe
			for row:=0; row<nSamples; row++ {
				c2 := xe - stateVals[iter].AtVec (row)
				rawC[col*nSamples+row] = c2*c2
			}
		}
		sg := proxRecur (distbs[iter], 
			mat.NewVecDense (nSamples, rawPsi),
			mat.NewDense (nSamples, nSamples, rawC),
			beta, h)

		// normalize
		Z := 0.0
		for col:=0; col<nSamples-1; col++ {
			dz := stateVals[iter+1].AtVec(col+1) - stateVals[iter+1].AtVec(col)
			Z += (sg.AtVec (col) + sg.AtVec(col+1)) * dz / 2.0
		}
		tmp1.Reset ()
		tmp1.ScaleVec (1.0/Z, sg)

		distbs[iter+1] = &tmp1
	}

	return stateVals, distbs
}

func plotDistb (state, distb mat.Vector, filename string, time, a, beta, mu0, stddev0 float64) {
	n := state.Len()
	if n != distb.Len() {
		panic ("Not accept different-length vectors.")
	}

	// plot proximal
	xys := make(plotter.XYs, n)
	for i := range(xys) {
		xys[i].X = state.AtVec(i)
		xys[i].Y = distb.AtVec(i)
	}
	sort.Slice (xys, func(i,j int) bool {
		return xys[i].X < xys[j].X
	})

	l1, err := plotter.NewLine (xys)
	if err != nil {
		panic (err)
	}
	l1.Color = &color.RGBA { 0x1f, 0x77, 0xb4, 0xff }

	// plot analytical
	xys2 := make (plotter.XYs, 353)
	idx := 0
	for t:=-4.0; t<=7.0; t+=0.03125 {
		xys2[idx].X = t
		tmp1 := t - mu0*math.Exp(-a*time)
		tmp2 := (stddev0*stddev0 - 1.0/(a*beta))*math.Exp(-2*a*time) + 1.0/(a*beta)
		xys2[idx].Y = math.Exp (-tmp1*tmp1/(2.0*tmp2)) / math.Sqrt(2.0*math.Pi*tmp2)
		idx ++
	}

	l2, err := plotter.NewLine (xys2)
	if err != nil {
		panic (err)
	}
	l2.Color = &color.RGBA{ 0xff, 0x7f, 0x0e, 0xff };

	// set plot
	p, err := plot.New()
	if err != nil {
		panic (err)
	}

	p.Add (l2)
	p.Add (l1)

	p.Legend.Add ("analytical", l2)
	p.Legend.Add ("proximal", l1)
	p.Legend.Top = true
	p.Legend.Left = true

	p.X.Min = -4.0
	p.X.Max = 7.0
	p.X.Label.Text = "x"
	p.Y.Min = 0.0
	p.Y.Max = 2.0
	p.Y.Label.Text = "p(x,t)"

	// save
	if err := p.Save (3*vg.Inch, 3*vg.Inch, filename+".png"); err != nil {
		panic (err)
	}
	if err := p.Save (3*vg.Inch, 3*vg.Inch, filename+".eps"); err != nil {
		panic (err)
	}
}

func plotMain (stateVals, distbs []mat.Vector, a, beta, h, mu0, stddev0 float64) {
	repTime := [...]int{0, 500, 1000, 2000, 3000, 4000}
	// repTime := [...]int{0, 10}
	for _, t := range (repTime) {
		filename := fmt.Sprintf ("./figs/ornstein_uhlenbeck_t_%04d", t)
		plotDistb (stateVals[t], distbs[t], filename, float64(t)*h, a, beta, mu0, stddev0)
	}
}


func main() {
	rnd = rand.New (rand.NewSource(1234))

	// set parameter
	mu0 := 5.0
	stddev0 := 0.2
	variance0 := stddev0 * stddev0
	nSamples := 400

	a := 1.0
	beta := 1.0
	h := 0.001

	// initialize
	rawInit := make ([]float64, nSamples)
	rawInitDistbs := make ([]float64, nSamples)
	for col:=0; col<nSamples; col++ {
		rv := stddev0 * rnd.NormFloat64() 
		rawInit[col] = rv + mu0
		rawInitDistbs[col] = math.Exp (-rv*rv/(2.0*variance0)) / math.Sqrt (2.0*math.Pi*variance0)
	}

	// calcuration
	resState, resDistbs := eulerMaruyama (
		mat.NewVecDense (nSamples, rawInit),
		mat.NewVecDense (nSamples, rawInitDistbs),
		a, beta, h)
	
	// plot
	plotMain (resState, resDistbs, a, beta, h, mu0, stddev0)
}

