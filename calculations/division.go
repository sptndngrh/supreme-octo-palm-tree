package calculations

func Divide(a, b float64) float64 {
    if b == 0 {
        return 0 // Handle division by zero
    }
    return a / b
}
