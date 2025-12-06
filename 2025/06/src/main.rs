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

#[derive(Debug)]
enum Op {
    Add,
    Multiply,
}

fn parse_input(input: &str) -> Option<(Vec<Vec<usize>>, Vec<Op>)> {
    let lines: Vec<&str> = input.trim().lines().map(|line| line).collect();
    let (ops_str, numbers_strs) = lines.split_last()?;

    let ops: Vec<Op> = ops_str
        .split_whitespace()
        .filter_map(|o| match o {
            "+" => Some(Op::Add),
            "*" => Some(Op::Multiply),
            _ => None,
        })
        .collect();

    let numbers: Vec<Vec<usize>> = numbers_strs
        .iter()
        .map(|line| {
            line.split_whitespace()
                .filter_map(|str| str.parse::<usize>().ok())
                .collect::<Vec<usize>>()
        })
        .collect();

    Some((numbers, ops))
}

fn part1(input: &str) -> String {
    let Some((numbers, ops)) = parse_input(input) else {
        return String::new();
    };

    let Some(column_results) = numbers.iter().cloned().reduce(|acc, row| {
        acc.iter()
            .zip(row.iter())
            .enumerate()
            .map(|(idx, (a, b))| match ops[idx] {
                Op::Add => a + b,
                Op::Multiply => a * b,
            })
            .collect()
    }) else {
        return String::new();
    };

    column_results.iter().sum::<usize>().to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "4277556");
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
