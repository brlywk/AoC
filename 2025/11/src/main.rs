use std::collections::HashMap;

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
struct Device {
    // NOTE: As our puzzle input is a &'static str (included in the program)
    // we can keep using &str's, b/c the references will be guaranteed to live
    // as long as the program itself
    label: &'static str,
    outputs: Vec<&'static str>,
}

impl Device {
    // note: as this is a full list of devices, we can parse single lines, non-recursive
    fn parse(s: &'static str) -> Option<Self> {
        let (label, rest) = s.split_once(':')?;
        let outputs = rest
            .split(' ')
            .filter(|s| !s.is_empty())
            .collect::<Vec<&str>>();

        Some(Self { label, outputs })
    }
}

fn parse_input(input: &'static str) -> Vec<Device> {
    input.trim().lines().filter_map(Device::parse).collect()
}

// part 1 should be a "simple" DFS with backtracking
fn find_all_paths(
    devices: &[Device],
    start: &'static str,
    end: &'static str,
) -> Vec<Vec<&'static str>> {
    // convert to lookup map for devices by label
    let device_map: HashMap<&str, &Device> = devices.iter().map(|d| (d.label, d)).collect();

    let mut all_paths = Vec::new();
    let mut stack = vec![(start, vec![start])];

    while let Some((current, path)) = stack.pop() {
        if current == end {
            all_paths.push(path);
            continue;
        }

        if let Some(device) = device_map.get(current) {
            for &output in &device.outputs {
                let mut new_path = path.clone();
                new_path.push(output);
                stack.push((output, new_path));
            }
        }
    }

    all_paths
}

const START: &str = "you";
const END: &str = "out";

fn part1(input: &'static str) -> String {
    find_all_paths(&parse_input(input), START, END)
        .iter()
        .len()
        .to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "5");
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
