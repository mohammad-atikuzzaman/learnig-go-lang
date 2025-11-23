package utils

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

func CalculateAverage(numbers []float64) float64 {
    if len(numbers) == 0 {
        return 0
    }
    
    total := 0.0
    for _, num := range numbers {
        total += num
    }
    
    return total / float64(len(numbers))
}