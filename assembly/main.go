package main

import "C"

//export SIMDSumArray
func SIMDSumArray(arr []float32) float32 {
    return simdSumArrayAsm(arr)
}

func simdSumArrayAsm(arr []float32) float32

func main() {}
