package main

import (
	"fmt"
	"sort"
)

type Building struct {
	X, Y, Profit int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	buildings := make([]Building, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &buildings[i].X, &buildings[i].Y, &buildings[i].Profit)
	}

	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i].X == buildings[j].X {
			return buildings[i].Y < buildings[j].Y
		}
		return buildings[i].X < buildings[j].X
	})

	dp := make([]int, n)
	dp[0] = buildings[0].Profit

	for i := 1; i < n; i++ {
		dp[i] = buildings[i].Profit
		for j := 0; j < i; j++ {
			if buildings[j].X < buildings[i].X && buildings[j].Y < buildings[i].Y ||
				buildings[j].X < buildings[i].X && buildings[j].Y > buildings[i].Y ||
				buildings[j].X > buildings[i].X && buildings[j].Y < buildings[i].Y ||
				buildings[j].X > buildings[i].X && buildings[j].Y > buildings[i].Y {
				dp[i] = max(dp[i], dp[j]+buildings[i].Profit)
			}
		}
	}

	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
