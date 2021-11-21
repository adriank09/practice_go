package main

func Minus(n1 int, n2 int) int {
	return n1 - n2
}

func Multiply(n1 int, n2 int) int {
	return n1 * n2
}

func Calculate(n1 int, n2 int, fn func(int, int) int) int {
	return fn(n1, n2)
}
