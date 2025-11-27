package main

import "core:fmt"
import "core:log"
import "core:strings"
import "core:testing"

TEST :: #load("./test.txt")
DATA :: #load("./data.txt")

main :: proc() {
	p1 := part1(DATA)
	fmt.println("Part 1:", p1)

	p2 := part2(DATA)
	fmt.println("Part 2:", p2)
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

SearchDirection :: enum {
	Forward,
	Backward,
	Up,
	Down,
	ToTopLeft,
	ToTopRight,
	ToBottomLeft,
	ToBottomRight,
}

@(rodata)
Directions := [SearchDirection][2]int {
	.Forward       = {0, 1},
	.Backward      = {0, -1},
	.Up            = {-1, 0},
	.Down          = {1, 0},
	.ToTopLeft     = {-1, -1},
	.ToTopRight    = {-1, 1},
	.ToBottomLeft  = {1, -1},
	.ToBottomRight = {1, -1},
}

// internal matrix type only supports 16x16 matrices :(
Matrix :: struct {
	data: [][]u8,
	rows: int,
	cols: int,
}

matrix_init :: proc(input: []u8) -> Matrix {
	rows := 0
	cols := 0

	it := string(input)
	for _ in strings.split_lines_iterator(&it) {
		rows += 1
	}

	data := make([][]u8, rows, context.temp_allocator)

	i := 0
	it = string(input)
	for line in strings.split_lines_iterator(&it) {
		data[i] = make([]u8, len(line), context.temp_allocator)
		copy(data[i], line)
		cols = max(cols, len(line))
		i += 1
	}

	return {data = data, rows = rows, cols = cols}
}

matrix_deinit :: proc(m: ^Matrix) {
	for row in m.data {
		delete(row, context.temp_allocator)
	}
	delete(m.data, context.temp_allocator)
}

matrix_get :: proc(m: Matrix, row, col: int) -> (u8, bool) {
	if row > len(m.data) || col > len(m.data[0]) do return 0, false
	return m.data[row][col], true
}

matrix_search :: proc(m: Matrix, word: string, row, col: int, dir: SearchDirection) -> bool {
	wl := len(word)
	if wl == 0 do return false

	dir_vec := Directions[dir]

	// first letter must match
	if c, ok := matrix_get(m, row, col); !ok || c != word[0] do return false

	// check bounds
	end_row := row + dir_vec.x * (wl - 1)
	end_col := col + dir_vec.y * (wl - 1)

	if end_row < 0 || end_row >= m.rows || end_col < 0 || end_col >= m.cols {
		return false
	}

	// check other letters in word
	for letter, idx in word {
		r := row + dir_vec.x * idx
		c := col + dir_vec.y * idx

		if c, ok := matrix_get(m, r, c); !ok || c != u8(letter) {
			return false
		}
	}

	return true
}

matrix_find_word :: proc(m: Matrix, word: string, row, col: int) -> int {
	count := 0

	for dir in SearchDirection {
		if matrix_search(m, word, row, col, dir) do count += 1
	}

	return count
}

matrix_find_word_all :: proc(m: Matrix, word: string) -> int {
	count := 0

	for row, r in m.data do for _, c in row {
		count += matrix_find_word(m, word, r, c)
	}

	return count
}


part1 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	m := matrix_init(input)
	defer matrix_deinit(&m)

	count := matrix_find_word_all(m, "XMAS")

	return fmt.aprint(count)
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "18"
	actual := part1(TEST)
	defer delete(actual)


	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

matrix_find_x_mas :: proc(m: Matrix) -> int {
	count := 0

	for row, r in m.data do for col, c in row {
		// look for 'A' as pivot
		if col != 'A' do continue

		// bounds check (if at borders there can be "no mas" :P)
		if c == 0 || c == len(row) - 1 || r == 0 || r == len(m.data) - 1 {
			continue
		}

		// check diagonals for M-S or S-M pairs
		// top-left <-> bottom-right
		tl_br_valid := (m.data[r - 1][c - 1] == 'M' && m.data[r + 1][c + 1] == 'S') || (m.data[r - 1][c - 1] == 'S' && m.data[r + 1][c + 1] == 'M')

		// top-right <-> bottom-left
		tr_bl_valid := (m.data[r - 1][c + 1] == 'M' && m.data[r + 1][c - 1] == 'S') || (m.data[r - 1][c + 1] == 'S' && m.data[r + 1][c - 1] == 'M')

		if tl_br_valid && tr_bl_valid do count += 1
	}

	return count
}

part2 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	m := matrix_init(input)
	defer matrix_deinit(&m)

	count := matrix_find_x_mas(m)

	return fmt.aprint(count)
}

@(test)
part2_test :: proc(t: ^testing.T) {
	expected := "9"
	actual := part2(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

