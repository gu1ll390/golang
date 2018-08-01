package main

import (
	"fmt"
	"strconv"
)

// Complete the queensAttack function below.
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {

	if n == 1 {
		return 0
	}

	m := make(map[string]bool)
	var i int32
	for i = 0; i < k; i++ {
		s := strconv.FormatInt(int64(obstacles[i][0]-1), 10) + "," + strconv.FormatInt(int64(obstacles[i][1]-1), 10)
		m[s] = true
	}

	dir := [][]int32{[]int32{-1, 0}, []int32{-1, 1}, []int32{0, 1}, []int32{1, 1}, []int32{1, 0}, []int32{1, -1}, []int32{0, -1}, []int32{-1, -1}}

	var moves int32
	for i := 0; i < 8; i++ {
		moves += queensAttackAux(n, k, r_q-1+dir[i][0], c_q-1+dir[i][1], obstacles, dir[i], m)
	}

	return moves
}

func queensAttackAux(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32, dir []int32, m map[string]bool) int32 {

	key := strconv.FormatInt(int64(r_q), 10) + "," + strconv.FormatInt(int64(c_q), 10)

	if r_q < 0 || r_q >= n || c_q < 0 || c_q >= n || m[key] {
		return 0
	}

	return 1 + queensAttackAux(n, k, r_q+dir[0], c_q+dir[1], obstacles, dir, m)
}

func main() {

	n := int32(10000)
	r_q := int32(5000)
	c_q := int32(5000)

	k := int32(10000)
	obstacles := make([][]int32, k)
	var i int32
	for i = 0; i < k; i++ {
		obstacles[i] = make([]int32, 2)
		obstacles[i][0] = 4000
		obstacles[i][1] = 5000
	}

	// k := int32(3)
	// obstacles := [][]int32{[]int32{5010, 0}, []int32{0, 5000}, []int32{3, 0}}

	result := queensAttack(n, k, r_q, c_q, obstacles)

	fmt.Print(result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
