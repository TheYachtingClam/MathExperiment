// My tools tt simulate Can't Stop the board game

package main

import (
	"fmt"
)

func cartN(a ...[]int) [][]int {
    c := 1
    for _, a := range a {
        c *= len(a)
    }
    if c == 0 {
        return nil
    }
    p := make([][]int, c)
    b := make([]int, c*len(a))
    n := make([]int, len(a))
    s := 0
    for i := range p {
        e := s + len(a)
        pi := b[s:e]
        p[i] = pi
        s = e
        for j, n := range n {
            pi[j] = a[j][n]
        }
        for j := len(n) - 1; j >= 0; j-- {
            n[j]++
            if n[j] < len(a[j]) {
                break
            }
            n[j] = 0
        }
    }
    return p
}

func possiblePairSums(a [][]int) map[int]int {
	m := make(map[int]int)
	for _, v := range a {
		if len(v) >= 4 {
			lm := make(map[int]int)
			lm[v[0]+v[1]]++
			lm[v[0]+v[2]]++
			lm[v[0]+v[3]]++	
			lm[v[1]+v[2]]++
			lm[v[1]+v[3]]++
			lm[v[2]+v[3]]++
			for k, _ := range lm {
				m[k]++
			}
		}
	}
	return m
}

func possiblePairSumsTargeted(a [][]int, b []int) map[int]int {
	m := make(map[int]int)
	for _, v := range a {
		if len(v) >= 4 {
			lm := make(map[int]int)
			lm[v[0]+v[1]]++
			lm[v[0]+v[2]]++
			lm[v[0]+v[3]]++	
			lm[v[1]+v[2]]++
			lm[v[1]+v[3]]++
			lm[v[2]+v[3]]++
			for k, _ := range lm {
				if b[0] == k || b[1] == k || b[2] == k {
					m[k]++
				}
			}
		}
	}
	return m
}

func chanceToSucceed(a [][]int, b []int) float32 {
	m := 0
	for _, v := range a {
		if len(v) >= 4 {
			lm := make(map[int]int)
			lm[v[0]+v[1]]++
			lm[v[0]+v[2]]++
			lm[v[0]+v[3]]++	
			lm[v[1]+v[2]]++
			lm[v[1]+v[3]]++
			lm[v[2]+v[3]]++
			for k, _ := range lm {
				if b[0] == k || b[1] == k || b[2] == k {
					m++
					break
				}
			}
		}
	}
	return float32(m) / 1296.0
}

func main() {
    orig := []int{1, 2, 3, 4, 5, 6}
	perm := cartN(orig, orig, orig, orig)
	vals := possiblePairSums(perm)
	v := possiblePairSumsTargeted(perm, []int{2, 12, 12})
    fmt.Println(perm)

	fmt.Println(len(perm))
	fmt.Println(vals)
	fmt.Println(v)
	fmt.Println(chanceToSucceed(perm, []int{3, 7, 4}))
}