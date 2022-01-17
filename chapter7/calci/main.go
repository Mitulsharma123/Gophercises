package calci

func Max(xs []float64) float64{
	max := xs[0]
	for _, v := range xs {
		if v > max {
			max = v 
		}
	}
	return max
}

func Min(xs []float64) float64 {
	min := xs[0]
	for _,v := range xs {
		if v < min {
			min = v
		}
	}
	return min 
}

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
