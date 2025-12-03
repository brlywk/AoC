#[allow(dead_code)]
const TEST: &str = include_str!("./test.txt");
const DATA: &str = include_str!("./data.txt");

fn main() {
    let p1 = part1(DATA);
    println!("Part 1: {p1}");

    // let p2 = part2(DATA);
    // println!("Part 2: {p2}");
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

fn part1(input: &str) -> String {
    // find highest number in line, then highest number in sub-slice from highest to end
    input
        .trim()
        .lines()
        .filter_map(|line| {
            let bytes = line.as_bytes();

            // we need to skip the last number b/c it cannot be the "first highest" number
            // BUT: we need to look in reverse to find the FIRST max in case of nasty gotchas like
            // 989 -> would take the second 9 as max (in correct order), but that would remove the
            //     8 as potential second highest
            let (index, highest) = bytes[..bytes.len() - 1]
                .iter()
                .enumerate()
                .rev()
                .max_by_key(|(_idx, byte)| *byte)?;
            let second_highest = line.as_bytes()[index + 1..].iter().max()?;

            let a = highest - b'0';
            let b = second_highest - b'0';

            format!("{a}{b}").parse::<usize>().ok()
        })
        .sum::<usize>()
        .to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "357");
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

// fn part2(input: &str) -> String {
//     let mut result: isize = 0;
//
//     result.to_string()
// }

// #[test]
// fn part2_test() {
//     assert_eq!(part2(TEST), "42");
// }
