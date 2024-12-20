package util

func Float64ToFloat32(slice []float64) []float32 {
	newSlice := make([]float32, len(slice))
	for i, v := range slice {
		newSlice[i] = float32(v)
	}
	return newSlice
}

func Float32ToFloat64(slice []float32) []float64 {
	newSlice := make([]float64, len(slice))
	for i, v := range slice {
		newSlice[i] = float64(v)
	}
	return newSlice
}
