use std::collections::{HashMap, HashSet};

#[allow(dead_code)]
const TEST: &str = include_str!("./test.txt");
const TEST2: &str = include_str!("./test2.txt");
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

const START_P1: &str = "you";
const END_P1: &str = "out";

fn part1(input: &'static str) -> String {
    find_all_paths(&parse_input(input), START_P1, END_P1)
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

const START_P2: &str = "svr";
const END_P2: &str = "out";
const FFT: &str = "fft";
const DAC: &str = "dac";

// try essentially the solution from part 1, but don't store, just count paths
fn count_paths(
    devices: &[Device],
    start: &str,
    end: &str,
    must_visit_1: &str,
    must_visit_2: &str,
) -> usize {
    // convert to lookup map for devices by label
    let device_map: HashMap<&str, &Device> = devices.iter().map(|d| (d.label, d)).collect();

    // to save some time, store a HashSet that just checks if a path visits both "must_visits"
    let mut count = 0;
    let mut stack = vec![(start, HashSet::from([start]), false, false)];

    while let Some((current, visited, has_1, has_2)) = stack.pop() {
        let new_has_1 = has_1 || current == must_visit_1;
        let new_has_2 = has_2 || current == must_visit_2;

        if current == end {
            if new_has_1 && new_has_2 {
                count += 1;
            }
            continue;
        }

        if let Some(device) = device_map.get(current) {
            for &output in &device.outputs {
                if !visited.contains(output) {
                    let mut new_visited = visited.clone();
                    new_visited.insert(output);
                    stack.push((output, new_visited, new_has_1, new_has_2));
                }
            }
        }
    }

    count
}

fn part2(input: &'static str) -> String {
    count_paths(&parse_input(input), START_P2, END_P2, FFT, DAC).to_string()
}

#[test]
fn part2_test() {
    assert_eq!(part2(TEST2), "2");
}
