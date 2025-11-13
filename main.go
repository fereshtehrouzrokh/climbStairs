package main

import "fmt"

// climbStairs returns the number of distinct ways to climb n steps
// when you can take 1 or 2 steps at a time.
// Recurrence: f(n) = f(n-1) + f(n-2), with f(0)=1, f(1)=1.
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	// f0 = f(0), f1 = f(1)
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}


func climbStairs(int n)int{

	if n<= 1 {
		return 1
	}
	f0:= 1
	f1:=1
	
	for i:= 2;i<=n;i++{
		f0 = f1
		f1= f0+f1
	}
	return f1
}

func main() {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1}, // 1 way: do nothing
		{1, 1}, // [1]
		{2, 2}, // [1+1], [2]
		{3, 3}, // [1+1+1], [1+2], [2+1]
		{4, 5},
		{5, 8},
		{10, 89},
	}

	for _, t := range tests {
		got := climbStairs(t.n)
		fmt.Printf("climbStairs(%d) = %d (want %d)\n", t.n, got, t.want)
	}
}
