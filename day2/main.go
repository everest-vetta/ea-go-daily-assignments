package main

func main() {
	qc := QuickCash{}
	withdrawRequest := []float64{500.0, 400.0, 200.0, 300.0, 100.0, 300.0}
	qc.GetCash(withdrawRequest)
}
