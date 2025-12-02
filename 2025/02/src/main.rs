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
fn repeated_twice(num: usize) -> bool {
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
                .map(|num| if repeated_twice(num) { num } else { 0 })
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

// Similar to repeated_twice(): treat num as string, and split it evenly up to its length
// and compare all parts with each other
fn is_invalid(num: usize) -> bool {
    let num_str = num.to_string();
    let l = num_str.len();

    // one of these MUST be true for the number to be invalid
    // num splits: l/2
    // one split is always at each digit
    // l = 2 -> 2/2 -> 1 split  -> 11     -> 1,1 (at l/2)
    // l = 4 -> 4/2 -> 2 splits -> 5050   -> 50,50 and 5,0,5,0 (at l/2)
    // l = 6 -> 6/2 -> 3 splits -> 121212 -> 121, 211 / 12,12,12 / 1,2,1,2,1,2
    // etc.
    (1..=l / 2)
        .filter(|pattern_length| l.is_multiple_of(*pattern_length))
        .any(|pattern_length| {
            let pattern = &num_str[..pattern_length];
            num_str == pattern.repeat(l / pattern_length)
        })
    // at this point I would like to thank clippy's pedantic settings for letting me learn
    // MANY new stdlib functions (e.g. is_multiple_of) ;)
}

fn part2(input: &str) -> String {
    parse_input(input.trim())
        .iter()
        .map(|range| {
            range
                .clone()
                .map(|num| if is_invalid(num) { num } else { 0 })
                .sum::<usize>()
        })
        .sum::<usize>()
        .to_string()
}

#[test]
fn part2_test() {
    assert_eq!(part2(TEST), "4174379265");
}
