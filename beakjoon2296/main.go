package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Building struct {
	x, y, c int
}

func comp(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 건물 수 받기
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	N, _ := strconv.Atoi(line)

	// 건물들 값 받기
	buildings := make([]Building, N)
	for i := 0; i < N; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		c, _ := strconv.Atoi(parts[2])
		buildings[i] = Building{x, y, c}
	}

	// x 기준 정렬
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i].x < buildings[j].x
	})

	// 동적 계획법
	dp := make([][2]int, N)
	for i := 0; i < N; i++ {
		dp[i][0] = buildings[i].c
		dp[i][1] = buildings[i].c
	}

	max := 0
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if buildings[j].y < buildings[i].y {
				dp[i][0] = comp(dp[i][0], dp[j][0]+buildings[i].c)
			}
			if buildings[j].y > buildings[i].y {
				dp[i][1] = comp(dp[i][1], dp[j][1]+buildings[i].c)
			}
		}
		max = comp(max, comp(dp[i][0], dp[i][1]))
	}

	fmt.Fprintln(writer, max)
}
