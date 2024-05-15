package main

import (
	"fmt"
	"sort"
)
//Building 구조체를 정의
type Building struct {
	X, Y, Profit int//건물 좌표와 이익
}

func main() {
	var n int
	fmt.Scanf("%d", &n)//건물의 수
        //건물 정보 슬라이스 생성, 입력
	buildings := make([]Building, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &buildings[i].X, &buildings[i].Y, &buildings[i].Profit)
	}
        //X좌표로 정렬
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i].X == buildings[j].X { // X좌표가 같으면
			return buildings[i].Y < buildings[j].Y //Y좌표로 오름차순 정렬
		}
		return buildings[i].X < buildings[j].X//X좌표로 오름차순 정렬
	})
        //dp 배열 초기화
	dp := make([]int, n)
	dp[0] = buildings[0].Profit//첫 번째 건물 이익으로 초기화

	// 각 건물에 대한 최대 이익 계산
	for i := 1; i < n; i++ {
		dp[i] = buildings[i].Profit // 현재 건물의 이익으로 초기화
		for j := 0; j < i; j++ {
			//현재 건물을 지으면서 이전 건물과 관계가 있는 경우를 고려하여 최대 이익 계산
			if buildings[j].X < buildings[i].X && buildings[j].Y < buildings[i].Y || //1사분면
				buildings[j].X < buildings[i].X && buildings[j].Y > buildings[i].Y || //2사분면
				buildings[j].X > buildings[i].X && buildings[j].Y < buildings[i].Y || //3사분면
				buildings[j].X > buildings[i].X && buildings[j].Y > buildings[i].Y {  //4사분면
				dp[i] = max(dp[i], dp[j]+buildings[i].Profit) //최대 이익 갱신
			}
		}
	}

	//최대 이익 계산
	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	fmt.Println(ans)//최대 이익 출력
}
// 두 정수 중 큰 값을 반환하는 함수
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
