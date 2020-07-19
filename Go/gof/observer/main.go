package main;

func main() {
	var generator *RandomNumberGenerator = NewRandomNumberGenerator ()
	var digitObserver Observer = &DigitObserver {}
	var graphObserver Observer = &GraphObserver {}
	generator.AddObserver (digitObserver)
	generator.AddObserver (graphObserver)
	generator.Execute ()
}
