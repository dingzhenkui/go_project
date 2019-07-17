package main

import "fmt"

func rectPorts(aa, bb float64) (float64, float64) {
	var area = aa * bb
	var area1 = (aa + bb) * 2
	return area, area1
}
func main() {
	cc, dd := rectPorts(10.1, 22)
	fmt.Printf("cc is %f,\ndd is %f", cc, dd)
}
