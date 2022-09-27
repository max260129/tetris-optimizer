package main

import "fmt"

type board [][]byte

func makeBoard(size int) board {
	res := make([][]byte, size)
	for i := range res {
		res[i] = make([]byte, size)
		for j := range res[i] {
			res[i][j] = dot
		}
	}
	return res
}

func (b board) print() {
	for i := range b {
		fmt.Println(string(b[i]))
	}
}

func (b board) check(i, j int, t tetrimino) bool {
	for l := range t {
		for x := range t[l] {
			if t[l][x] == hashTag {
				if i+l >= len(b) || j+x >= len(b) {
					return false
				}
				if b[i+l][j+x] != dot {
					return false
				}
			}
		}
	}
	return true
}

func (b board) put(i, j, idx int, t tetrimino) {
	for y := range t {
		for x := range t[y] {
			if t[y][x] == hashTag {
				b[i+y][j+x] = byte(idx + 'A')
			}
		}
	}
}

func (b board) remove(i, j int, t tetrimino) {
	for y := range t {
		for x := range t[y] {
			if t[y][x] == hashTag {
				b[i+y][j+x] = dot
			}
		}
	}
}