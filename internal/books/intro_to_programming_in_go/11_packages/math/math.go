package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

// Finds the minimum value in a slice of numbers
func Min(xs []float64) float64 {
    minimum := xs[0]
    for _, x := range xs {
        if x < minimum {
            minimum = x
        }
    }
    return minimum
}

// Finds the maximum value in a slice of numbers
func Max(xs []float64) float64 {
    maximum := xs[0]
    for _, x := range xs {
        if x > maximum {
            maximum = x
        }
    }
    return maximum
}
