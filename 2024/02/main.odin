package main

import "core:fmt"
import "core:log"
import "core:strconv"
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

parse_input :: proc(input: []u8, allocator := context.allocator) -> [dynamic][dynamic]int {
	levels := make([dynamic][dynamic]int, allocator)

	it := string(input)
	for line in strings.split_lines_iterator(&it) {
		line_arr := make([dynamic]int, allocator)

		ll := line
		for l in strings.split_iterator(&ll, " ") {
			num, _ := strconv.parse_int(l)
			append(&line_arr, num)
		}

		append(&levels, line_arr)
	}

	return levels
}

Direction :: enum {
	None,
	Increasing,
	Decreasing,
}

level_safe :: proc(level: [dynamic]int) -> bool {
	a := level[0:len(level) - 1]
	b := level[1:len(level)]
	dir := Direction.None

	s := soa_zip(a = a, b = b)

	for n in s {
		// check if direction is steady
		current_dir := n.a > n.b ? Direction.Increasing : Direction.Decreasing

		if dir == .None do dir = current_dir
		if dir != current_dir do return false
		dir = current_dir

		// check if difference is within range 1 < diff < 3
		diff := abs(n.a - n.b)
		if diff < 1 || diff > 3 do return false
	}

	return true
}

part1 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	safe_levels := 0
	levels := parse_input(input, context.temp_allocator)

	for level in levels {
		if level_safe(level) do safe_levels += 1
	}

	return fmt.aprint(safe_levels)
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "2"
	actual := part1(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

part2 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	safe_levels := 0
	levels := parse_input(input, context.temp_allocator)

	for level in levels {
		is_safe := false

		for i in 0 ..< len(level) {
			perm := make([dynamic]int, len(level), context.temp_allocator)
			copy(perm[:], level[:])

			ordered_remove(&perm, i)

			if level_safe(perm) {
				is_safe = true
				break
			}
		}

		if is_safe do safe_levels += 1
	}

	return fmt.aprint(safe_levels)
}

@(test)
part2_test :: proc(t: ^testing.T) {
	expected := "4"
	actual := part2(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

