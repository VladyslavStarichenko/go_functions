package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// Define the range of values to search for roots
	a, b := 0.0, 4.0
	// Define the desired accuracy
	e := 1e-5

	// Create channels for sending intervals and receiving roots
	intervalChan := make(chan Interval)
	rootChan := make(chan float64)

	// Start a goroutine to find intervals using the Scanning method
	go scanIntervals(a, b, e, intervalChan)

	// Start multiple goroutines to find roots using different methods
	numMethods := 2
	for i := 0; i < numMethods; i++ {
		go findRoots(a, b, e, intervalChan, rootChan)
	}

	// Collect roots from all goroutines
	numRoots := 0
	var roots []float64
	for r := range rootChan {
		roots = append(roots, r)
		numRoots++
		if numRoots == numMethods {
			// All roots have been found, so stop receiving from the channel
			break
		}
	}

	// Write the roots to a file
	outFile, err := os.Create("roots.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	for _, r := range roots {
		fmt.Fprintf(outFile, "%f\n", r)
	}
}

// Define the equation to be solved
func f(x float64) float64 {
	return (x + 5) * (x - 1) * (x - 3)
}

// Define the derivative of the equation to be solved
func df(x float64) float64 {
	return 3*x*x - 9*x - 4
}

// Interval represents a closed interval [a, b]
type Interval struct {
	a, b float64
}

// Find intervals for processing using Scanning method
func scanIntervals(a, b, e float64, c chan Interval) {
	// Divide the range [a, b] into subintervals of width e
	for x := a; x < b; x += e {
		// Check if f(x) and f(x+e) have opposite signs
		if f(x)*f(x+e) < 0 {
			// If so, add the subinterval [x, x+e] to the channel
			c <- Interval{x, x + e}
		}
	}
	// Close the channel once all intervals have been sent
	close(c)
}

// Find roots of equation using bracketing methods
func findRoots(a, b, e float64, c chan Interval, r chan float64) {
	// Process each interval sent through the channel
	for i := range c {
		// Use the bisection method to find the root
		root := bisection(i.a, i.b, e)
		// Send the root back through the channel
		r <- root
	}
}

// Bisection method for finding the root of an equation within an interval
func bisection(a, b, e float64) float64 {
	fa, fb := f(a), f(b)
	// Check that the function values at the endpoints have opposite signs
	if fa*fb > 0 {
		panic("no root in interval")
	}
	// Keep dividing the interval in half until the width is less than e
	for math.Abs(b-a) > e {
		c := (a + b) / 2
		fc := f(c)
		if fa*fc < 0 {
			b, fb = c, fc
		} else {
			a, fa = c, fc
		}
	}
	// Return the midpoint of the final interval as the root
	return (a + b) / 2
}

// Find roots of equation using the secant method
func secant(a, b, e float64, r chan float64) {
	// Perform iterations until the width of the interval is less than e
	for math.Abs(b-a) > e {
		// Compute the next value in the sequence
		c := b - f(b)*(b-a)/(f(b)-f(a))
		// Move the endpoints of the interval
		a, b = b, c
	}
	// Send the final estimate of the root back through the channel
	r <- b
}
