#[allow(dead_code)]
const TEST: &str = include_str!("./test.txt");
const DATA: &str = include_str!("./data.txt");

fn main() {
    let p1 = part1(DATA);
    println!("Part 1: {p1}");

    let p2 = part2(DATA);
    println!("Part 2: {p2}");
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

fn parse_input(input: &str) -> Vec<std::ops::RangeInclusive<usize>> {
    input
        .split(',')
        .filter_map(|str| {
            let (str_a, str_b) = str.split_once('-')?;
            let a = str_a.parse::<usize>().ok()?;
            let b = str_b.parse::<usize>().ok()?;

            Some(a..=b)
        })
        .collect()
}

// There might be some mathematical magic here, but the easiest way to check twice
// repeating patterns is probably just "treat number as string, split in half, check equality" :D
//
// Could also be a trait, but that's probably a lot of boilerplate for little practical value ;)
fn is_repeated(num: usize) -> bool {
    let num_str = num.to_string();
    let (a, b) = num_str.split_at(num_str.len() / 2);

    // if a == b {
    //     println!("invalid id: {num}");
    // }

    a == b
}

fn part1(input: &str) -> String {
    parse_input(input.trim())
        .iter()
        .map(|range| {
            range
                .clone()
                .map(|num| if is_repeated(num) { num } else { 0 })
                .sum::<usize>()
        })
        .sum::<usize>()
        .to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "1227775554");
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

fn part2(input: &str) -> String {
    let mut result: isize = 0;

    result.to_string()
}

#[test]
fn part2_test() {
    assert_eq!(part2(TEST), "42");
}
