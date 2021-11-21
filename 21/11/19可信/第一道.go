package main

import (
	"fmt"
	"math"
)

func payByPromotion(prices []int, promotions [][]int, discounts []int, order []int) int {
	var maxPrice, maxPromotions int
	mpPrice := make(map[int]int, len(prices))
	mpOrder := make(map[int]int, len(order)/2)
	for i, price := range prices {
		mpPrice[i] = price
	}

	for i := 0; i < len(order); i += 2 {
		maxPrice += mpPrice[order[i]] * order[i+1]
		mpOrder[order[i]] = order[i+1]
	}

	// 贪心思想，根据优惠金额和达到优惠的总金额进行比较
	promotionsPrice := make([]int, len(discounts))
	for i, promotion := range promotions {
		for j := 0; j < len(promotion); j += 2 {
			promotionsPrice[i] = mpPrice[promotion[j]] * promotion[j+1]
		}
	}

	// 优惠力度
	promotionsLid := make([]float64, len(discounts))
	for i := 0; i < len(discounts); i++ {
		promotionsLid[i] = float64(promotionsPrice[i] * 1.0 / discounts[i])
	}

	// 找到位置
	mpPro := make([]int, len(discounts))
	for i := 0; i < len(discounts); i++ {
		mpPro[i] = findMin(promotionsLid)
		promotionsLid[mpPro[i]] = math.MaxFloat64
	}

	// 贪心求优惠的钱
	for i := 0; i < len(mpPro); i++ {
		site := mpPro[i]
		// 先判断够不够,记录次数
		isEnough := true
		time := 0
		tempOrder := mpOrder
		for isEnough {
			for j := 0; j < len(promotions[site]); j += 2 {
				if tempOrder[promotions[site][j]] < promotions[site][j+1] {
					isEnough = false
					break
				}
				tempOrder[promotions[site][j]] -= promotions[site][j+1]
			}
			if !isEnough {
				break
			}
			time++
		}

		// 得到次数，求优惠的钱
		maxPromotions += time * discounts[site]
		// 把order减去已经优惠完的数量
		for t := 0; t < time; t++ {
			for j := 0; j < len(promotions[site]); j += 2 {
				mpOrder[promotions[site][j]] -= promotions[site][j+1]
			}
		}
	}

	return maxPrice - maxPromotions
}

func findMin(lid []float64) int {
	var minSite int
	cnt := lid[0]
	for i := 0; i < len(lid); i++ {
		if lid[i] < cnt {
			cnt = lid[i]
			minSite = i
		}
	}

	return minSite
}

func main() {
	prices := []int{10, 5, 8, 8, 6, 3}
	promotions := [][]int{{0, 7}, {4, 10}, {2, 6, 1, 9}, {5, 2}}
	discounts := []int{4, 3, 5, 1}
	order := []int{2, 17, 3, 10, 1, 27, 5, 2, 4, 9}
	fmt.Println(payByPromotion(prices, promotions, discounts, order))

}
