package main

import "core:fmt"
import "core:strconv"
import "core:testing"
import "core:unicode"

TEST :: #load("./test.txt")
DATA :: #load("./data.txt")

main :: proc() {
	p1 := part1(DATA)
	fmt.println("Part 1:", p1)

	p2 := part1(DATA, true)
	fmt.println("Part 2:", p2)
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////


parse_num :: proc(input: string, idx: ^int) -> (int, bool) {
	i := idx^
	if i >= len(input) do return 0, false

	num_str := make([dynamic]u8, context.temp_allocator)

	for i < len(input) && unicode.is_digit(rune(input[i])) {
		append(&num_str, input[i])
		i += 1
	}

	idx^ = i
	num, _ := strconv.parse_int(string(num_str[:]))

	return num, true
}

matches_at :: proc(input: string, pos: int, pattern: string) -> bool {
	if (pos + len(pattern)) > len(input) do return false
	return input[pos:pos + len(pattern)] == pattern
}

Mul :: [2]int
parse_muls :: proc(input: string, handle_conditionals := false) -> [dynamic]Mul {
	muls := make([dynamic]Mul, context.temp_allocator)

	op_mul :: "mul("
	op_do :: "do()"
	op_dont :: "don't()"

	muls_enabled := true

	i := 0
	for i < len(input) {
		// check for "don't"
		if handle_conditionals && matches_at(input, i, op_dont) {
			muls_enabled = false
			i += len(op_dont)
			continue
		}

		// check for "do"
		if handle_conditionals && matches_at(input, i, op_do) {
			muls_enabled = true
			i += len(op_do)
			continue
		}

		// look for "mul("
		if matches_at(input, i, op_mul) {
			// move past mul
			i += len(op_mul)

			// parse first number "a": mul(a,b)
			a, _ := parse_num(input, &i)

			// expect a comma
			if i >= len(input) || input[i] != ',' do continue
			i += 1

			// parse second number "b": mul(a,b)
			b, _ := parse_num(input, &i)

			// expect closing ")"
			if i >= len(input) || input[i] != ')' do continue
			i += 1

			// if everything worked, we have a mul!
			if !handle_conditionals || muls_enabled {
				append(&muls, Mul{a, b})
			}
		} else {
			i += 1
		}
	}

	return muls
}

part1 :: proc(input: []u8, handle_conditionals := false) -> string {
	defer free_all(context.temp_allocator)

	sum := 0
	muls := parse_muls(string(input), handle_conditionals)

	for mul in muls {
		sum += mul.x * mul.y
	}

	return fmt.aprint(sum)
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "161"
	actual := part1(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////


@(test)
part2_test :: proc(t: ^testing.T) {
	new_test_str := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	new_test := transmute([]u8)new_test_str

	expected := "48"
	actual := part1(new_test, true)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

