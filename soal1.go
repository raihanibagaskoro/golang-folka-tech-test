package main

import (
	"fmt"
	"math"
)

func MaxProfit(prices []int, k int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	buy := make([]int, k+1)
	sell := make([]int, k+1)

	for i := 0; i <= k; i++ {
		buy[i] = math.MinInt32
		sell[i] = 0
	}

	for _, price := range prices {
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sell[j-1]-price)
			sell[j] = max(sell[j], buy[j]+price)
		}
	}

	return sell[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MaxProfit([]int{4, 11, 2, 20, 59, 80}, 2))
}