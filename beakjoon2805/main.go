package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 값 받기 (N, M)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	// 나무들 높이 받기
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	treeStrs := strings.Split(line, " ")
	trees := make([]int, N)
	for i := 0; i < N; i++ {
		trees[i], _ = strconv.Atoi(treeStrs[i])
	}
	sort.Ints(trees)

	// 이진탐색
	var low, height int
	high := trees[N-1]

	for low <= high {
		mid := (low + high) / 2
		var treeSum int
		for _, h := range trees {
			if h > mid {
				treeSum += h - mid
			}
		}
		if treeSum >= M {
			height = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Fprintln(writer, height)
}
