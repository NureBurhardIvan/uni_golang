package main

func IsGeometricProgression(N int) bool {
	a := N / 1000       
	b := (N / 100) % 10 
	c := (N / 10) % 10  
	d := N % 10         

	return (b*b == a*c) && (c*c == b*d)
}