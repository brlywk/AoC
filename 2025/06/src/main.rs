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
    let lines: Vec<&str> = input.trim().lines().collect();
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

fn really_annoying_way_to_parse_input(input: &str) -> Option<(Vec<Vec<usize>>, Vec<Op>)> {
    // idea to parse the input:
    // - each "virtual column" of left- or right-aligned numbers always
    //   starts at the index the operator is in
    // - the column before the operator is always a "spacer column" (only whitespaces)
    // -> go through the input column-wise, and parse each column as a single number (ignoring
    //    whitespaces)

    // AoC this year is OFFICIALLY EVIL! Not trimming the input usually screws you sideways,
    // but for this part trimming will produce the wrong result. EVIL!
    //
    // Scratch that: SUPER EVIL. Even the devil is green with envy right now, b/c not trimming
    // messes with the test, but trimming messes with the actual input... you really need to
    // explicitely only trim on newlines 😱
    let lines: Vec<&str> = input
        .trim_matches(|c| c == '\r' || c == '\n')
        .lines()
        .collect();
    let (ops_str, numbers_strs) = lines.split_last()?;

    let mut op_indices: Vec<usize> = Vec::with_capacity(ops_str.len());
    let mut ops: Vec<Op> = Vec::with_capacity(ops_str.len());

    // find the ops and there index in the string
    for (idx, c) in ops_str.chars().enumerate() {
        match c {
            '+' => {
                ops.push(Op::Add);
                op_indices.push(idx);
            }
            '*' => {
                ops.push(Op::Multiply);
                op_indices.push(idx);
            }
            _ => {}
        }
    }

    // now, split each row at the given indices to get the "virtual columns"
    // keep them as a string so we still keep the whitespaces
    let blocks: Vec<Vec<&str>> = (0..op_indices.len())
        .map(|idx| {
            let start = op_indices[idx];
            let end = op_indices.get(idx + 1).copied();

            numbers_strs
                .iter()
                // NOTE to self: filter_map makes it a lot more concise to parse lines in,
                // but considering the bottomless pit of evil of this years part2, it can also
                // be a major source of "WTF" if it silently drops parts of the input...
                .filter_map(|line| {
                    // get part of slice until the next col index, or to the end of the line
                    let slice = match end {
                        Some(col_end) => line.get(start..col_end)?,
                        None => line.get(start..)?,
                    };

                    // NOTE: We keep the "spacing" whitespace at the end, b/c that way we don't
                    //       handle the last column differntly and start calculating how "long"
                    //       the slice is supposed to be...

                    Some(slice)
                })
                .collect()
        })
        .collect();

    // now we can finally parse the numbers in each block by going through them
    // column by column
    let numbers: Vec<Vec<usize>> = blocks
        .iter()
        .map(|block| {
            let width = block.first().map_or(0, |s| s.len());

            (0..width)
                .filter_map(|col_idx| {
                    // use "None" so we can also filter out all Zeros created in the last step
                    block.iter().fold(None, |acc, row| {
                        let byte = row.as_bytes().get(col_idx).copied().unwrap_or(b' ');

                        if byte.is_ascii_digit() {
                            let digit = (byte - b'0') as usize;
                            Some(acc.unwrap_or(0) * 10 + digit)
                        } else {
                            acc
                        }
                    })
                })
                .collect()
        })
        .collect();

    Some((numbers, ops))
}

fn part2(input: &str) -> String {
    let Some((numbers, ops)) = really_annoying_way_to_parse_input(input) else {
        return String::new();
    };

    numbers
        .iter()
        .enumerate()
        .map(|(idx, block)| match ops[idx] {
            Op::Add => block.iter().sum(),
            Op::Multiply => block.iter().product::<usize>(),
        })
        .sum::<usize>()
        .to_string()
}

#[test]
fn part2_test() {
    assert_eq!(part2(TEST), "3263827");
}
