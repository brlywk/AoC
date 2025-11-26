package main

import "core:fmt"
import "core:slice"
import "core:sort"
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

number_lists :: proc(input: []u8) -> (left, right: [dynamic]int) {
	left = make([dynamic]int, context.temp_allocator)
	right = make([dynamic]int, context.temp_allocator)

	it := string(input)
	for line in strings.split_lines_iterator(&it) {
		// let's assume AoC input is always well formatted and we don't need to check
		// that we get the correct number of elements on split :)
		nums := strings.split(line, "   ", context.temp_allocator)
		l, _ := strconv.parse_int(nums[0])
		r, _ := strconv.parse_int(nums[1])

		append(&left, l)
		append(&right, r)
	}

	return
}

part1 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	left, right := number_lists(input)
	slice.sort(left[:])
	slice.sort(right[:])

	sum := 0
	for l, i in left {
		sum += abs(l - right[i])
	}

	return fmt.aprint(sum)
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "11"
	actual := part1(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

part2 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	left, right := number_lists(input)

	sum := 0
	for l in left {
		count := 0

		for r in right do if l == r {
			count += 1
		}

		sum += l * count
	}

	return fmt.aprint(sum)
}

@(test)
part2_test :: proc(t: ^testing.T) {
	expected := "31"
	actual := part2(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

