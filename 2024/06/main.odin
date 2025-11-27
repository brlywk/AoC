package main

import "core:fmt"
import "core:log"
import "core:slice"
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

Direction :: enum {
	Up,
	Down,
	Left,
	Right,
}

@(rodata)
Direction_Vec := [Direction][2]int {
	.Up    = {-1, 0},
	.Down  = {1, 0},
	.Left  = {0, -1},
	.Right = {0, 1},
}

WALL :: '#'
GUARD_START :: '^'

turn :: proc(dir: Direction) -> Direction {
	switch dir {
	case .Up:
		return .Right
	case .Right:
		return .Down
	case .Down:
		return .Left
	case .Left:
		return .Up
	}

	return dir
}

Grid :: [dynamic][]u8

input_parse :: proc(input: []u8) -> Grid {
	grid := make(Grid, context.temp_allocator)

	it := string(input)
	for line in strings.split_lines_iterator(&it) {
		append(&grid, transmute([]u8)line)
	}

	return grid
}

move :: proc(from: [2]int, dir: Direction) -> ([2]int, bool) {
	// vector math ftw!
	to := from + Direction_Vec[dir]
	return to, to.x >= 0 && to.y >= 0
}

grid_field :: proc(grid: Grid, at: [2]int) -> (u8, bool) {
	rows, cols := len(grid), len(grid[0])

	if at.x >= 0 && at.x < rows && at.y >= 0 && at.y < cols {
		return grid[at.x][at.y], true
	}

	return 0, false
}

find_guard :: proc(grid: Grid) -> [2]int {
	for row, r in grid {
		c := slice.linear_search(row[:], GUARD_START) or_continue
		return {r, c}
	}

	return {0, 0}
}

part1 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	grid := input_parse(input)
	visited := make(map[[2]int]struct{}, context.temp_allocator)

	guard_pos := find_guard(grid)
	guard_dir := Direction.Up
	visited[guard_pos] = {}

	// begin highly sophisticated movement simulation
	for {
		// what's the next field we would move to?
		move_next := move(guard_pos, guard_dir) or_break
		field_value := grid_field(grid, move_next) or_break

		// if occupied turn...
		if field_value == WALL {
			guard_dir = turn(guard_dir)
		} else {
			// ...or move there if free
			guard_pos = move_next
			visited[guard_pos] = {}
		}
	}

	return fmt.aprint(len(visited))
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "41"
	actual := part1(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

State :: struct {
	row: int,
	col: int,
	dir: Direction,
}

state_new_row_col :: proc(row, col: int, dir: Direction) -> State {
	return {row = row, col = col, dir = dir}
}

state_new_pos :: proc(pos: [2]int, dir: Direction) -> State {
	return {row = pos.x, col = pos.y, dir = dir}
}

state_new :: proc {
	state_new_row_col,
	state_new_pos,
}

check_loop :: proc(grid: Grid, start_pos, obstacle_pos: [2]int, start_dir: Direction) -> bool {
	states := make(map[State]struct{}, context.temp_allocator)

	guard_pos := start_pos
	guard_dir := start_dir

	for {
		// if we visited the same position twice, we be loopin'
		state := state_new(guard_pos, guard_dir)
		if state in states do return true

		states[state] = {}

		move_next := move(guard_pos, guard_dir) or_return
		field_value := grid_field(grid, move_next) or_return

		if move_next.x == obstacle_pos.x && move_next.y == obstacle_pos.y {
			field_value = WALL
		}

		if field_value == WALL {
			guard_dir = turn(guard_dir)
		} else {
			guard_pos = move_next
		}
	}
}

part2 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	grid := input_parse(input)
	visited := make(map[[2]int]struct{}, context.temp_allocator)

	guard_pos := find_guard(grid)
	guard_dir := Direction.Up
	start_pos := guard_pos
	start_dir := guard_dir
	visited[guard_pos] = {}

	// begin highly sophisticated movement simulation
	for {
		// what's the next field we would move to?
		move_next := move(guard_pos, guard_dir) or_break
		field_value := grid_field(grid, move_next) or_break

		// if occupied turn...
		if field_value == WALL {
			guard_dir = turn(guard_dir)
		} else {
			// ...or move there if free
			guard_pos = move_next
			visited[guard_pos] = {}
		}
	}

	// check all visited locations to see if putting an obstacle
	// there would cause a loop
	loop_count := 0
	checked_count := 0

	for v in visited {
		if v.x == start_pos.x && v.y == start_pos.y do continue

		checked_count += 1
		if check_loop(grid, start_pos, v, start_dir) {
			loop_count += 1
		}
	}

	return fmt.aprint(loop_count)
}

@(test)
part2_test :: proc(t: ^testing.T) {
	expected := "6"
	actual := part2(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

