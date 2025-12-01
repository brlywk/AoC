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

fn parse_input(input: &str) -> Vec<(&str, isize)> {
    input
        .lines()
        .filter_map(|line| {
            let (dir, num_str) = line.split_at(1);
            let n = num_str.parse::<isize>().ok()?;

            Some((dir, n))
        })
        .collect()
}

fn part1(input: &str) -> String {
    let mut result: isize = 0;
    let mut dial: isize = 50;

    let dial_movements = parse_input(input);
    // println!("movements: {dial_movements:?}");

    for (dir, n) in dial_movements {
        dial = match dir {
            "L" => (dial - n).rem_euclid(100),
            "R" => (dial + n).rem_euclid(100),
            _ => continue,
        };

        result += isize::from(dial == 0);

        // println!("result for: {dir}:{n} -> {dial} (result: {result})");
    }

    result.to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "3");
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

fn part2(input: &str) -> String {
    let mut result: isize = 0;
    let mut dial: isize = 50;

    let dial_movements = parse_input(input);

    for (dir, n) in dial_movements {
        // previous dial position
        let old_dial = dial;

        // new dial position
        dial = match dir {
            "L" => (dial - n).rem_euclid(100),
            "R" => (dial + n).rem_euclid(100),
            _ => continue,
        };

        // count if we crossed 0 (can cross 0 multiple times per movement)
        result += match dir {
            // to the left:
            // check how often we do a full rotation and if that full rotation crosses or lands on
            // 0 (but we didn't start at 0)
            "L" => n / 100 + isize::from(n % 100 >= old_dial && old_dial > 0),
            // to the right:
            // check how often we do a full rotation, but if we don't do a full rotation and land
            // exactly on 0 again, we don't want to count that, UNLESS we started at 0...
            "R" => {
                let sum = old_dial + n;
                let full_rotation = sum / 100;
                (old_dial + n) / 100
                    + isize::from(full_rotation == 0 && sum % 100 == 0 && old_dial > 0)
            }
            _ => 0,
        };

        // println!("result for: {dir}:{n} -> {dial} (crossed 0: {result})");
    }

    result.to_string()
}

#[test]
fn part2_test() {
    assert_eq!(part2(TEST), "6");
}
