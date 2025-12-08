use std::ops::{Deref, DerefMut};

#[allow(dead_code)]
const TEST: &str = include_str!("./test.txt");
const DATA: &str = include_str!("./data.txt");

fn main() {
    let p1 = part1(DATA, 1000);
    println!("Part 1: {p1}");

    let p2 = part2(DATA);
    println!("Part 2: {p2}");
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

// note: potential precision loss and need for upcasting,
//       but also good enough for this puzzle...
#[derive(Debug, PartialEq, Eq, Clone, Copy, Hash)]
struct Vec3 {
    x: i32,
    y: i32,
    z: i32,
}

impl Vec3 {
    fn zero() -> Self {
        Self { x: 0, y: 0, z: 0 }
    }

    fn parse(input: &str) -> Option<Self> {
        let mut nums = input.trim().split(',');

        Some(Self {
            x: nums.next()?.parse().ok()?,
            y: nums.next()?.parse().ok()?,
            z: nums.next()?.parse().ok()?,
        })
    }

    fn euclid_distance(&self, other: &Self) -> f64 {
        let dx = i64::from(self.x - other.x);
        let dy = i64::from(self.y - other.y);
        let dz = i64::from(self.z - other.z);

        // we need to upcast to prevent overflows, so just take the potential
        // precision loss and move on...
        #[allow(clippy::cast_precision_loss)]
        ((dx * dx + dy * dy + dz * dz) as f64).sqrt()
    }
}

#[derive(Debug)]
struct Circuit(Vec<Vec3>);

impl Deref for Circuit {
    type Target = Vec<Vec3>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for Circuit {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}

impl Circuit {
    fn new() -> Self {
        Self(Vec::new())
    }
}

fn parse_input(input: &str) -> Vec<Vec3> {
    input
        .trim()
        .lines()
        // note: this might gobble up invalid lines, but in general we can assume
        //       AoC input to be well formatted (albeit evil at times)...
        .filter_map(Vec3::parse)
        .collect()
}

fn part1(input: &str, n: usize) -> String {
    let junctions = parse_input(input);
    let mut circuits: Vec<Circuit> = Vec::new();

    let mut distances: Vec<(Vec3, Vec3, f64)> = Vec::new();

    // calculate the distance between each junction
    for i in 0..junctions.len() {
        for j in (i + 1)..junctions.len() {
            let dist = junctions[i].euclid_distance(&junctions[j]);
            distances.push((junctions[i], junctions[j], dist));
        }
    }

    // sort by shortest distance
    distances.sort_by(|a, b| a.2.partial_cmp(&b.2).unwrap());

    // for the n shortest distances...
    for (a, b, _dist) in distances.iter().take(n) {
        // check if a and b are already in a circuit
        let a_idx = circuits.iter().position(|c| c.contains(a));
        let b_idx = circuits.iter().position(|c| c.contains(b));

        match (a_idx, b_idx) {
            // both not in circuit: create new
            (None, None) => {
                let mut circuit = Circuit::new();
                circuit.push(*a);
                circuit.push(*b);
                circuits.push(circuit);
            }
            // a or b in circuit: add the other one
            (Some(idx), None) => {
                circuits[idx].push(*b);
            }
            (None, Some(idx)) => {
                circuits[idx].push(*a);
            }
            // both in different circuits: merge
            (Some(a_idx), Some(b_idx)) => {
                if a_idx != b_idx {
                    let circuit_b = circuits.remove(b_idx.max(a_idx));
                    let circuit_a_idx = a_idx.min(b_idx);
                    circuits[circuit_a_idx].extend(circuit_b.iter());
                }
            }
        }
    }

    // sort circuits by circuit length
    circuits.sort_by_key(|c| std::cmp::Reverse(c.len()));

    circuits
        .iter()
        .take(3)
        .map(|c| c.len())
        .product::<usize>()
        .to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST, 10), "40");
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

fn part2(input: &str) -> String {
    let junctions = parse_input(input);
    let n = junctions.len();
    let mut circuits: Vec<Circuit> = Vec::new();

    let mut distances: Vec<(Vec3, Vec3, f64)> = Vec::new();

    // calculate the distance between each junction
    for i in 0..junctions.len() {
        for j in (i + 1)..junctions.len() {
            let dist = junctions[i].euclid_distance(&junctions[j]);
            distances.push((junctions[i], junctions[j], dist));
        }
    }

    // sort by shortest distance
    distances.sort_by(|a, b| a.2.partial_cmp(&b.2).unwrap());

    // track which is the last connection
    let mut last_connected = (Vec3::zero(), Vec3::zero());
    let mut num_circuits = n;

    // this time do this for all distances, not just n
    for (a, b, _dist) in &distances {
        // check if a and b are already in a circuit
        let a_idx = circuits.iter().position(|c| c.contains(a));
        let b_idx = circuits.iter().position(|c| c.contains(b));

        match (a_idx, b_idx) {
            // both not in circuit: create new
            (None, None) => {
                let mut circuit = Circuit::new();
                circuit.push(*a);
                circuit.push(*b);
                circuits.push(circuit);
            }
            // a or b in circuit: add the other one
            (Some(idx), None) => {
                circuits[idx].push(*b);
            }
            (None, Some(idx)) => {
                circuits[idx].push(*a);
            }
            // both in different circuits: merge
            (Some(a_idx), Some(b_idx)) => {
                if a_idx == b_idx {
                    continue;
                }

                let circuit_b = circuits.remove(b_idx.max(a_idx));
                let circuit_a_idx = a_idx.min(b_idx);
                circuits[circuit_a_idx].extend(circuit_b.iter());
            }
        }

        // we did some connecting
        last_connected = (*a, *b);
        num_circuits -= 1;

        // if there is only 1 circuit left, we're done
        if num_circuits == 1 {
            break;
        }
    }

    // stupid integer overlow being all stupid and panicking the program again...
    (i64::from(last_connected.0.x) * i64::from(last_connected.1.x)).to_string()
}

#[test]
fn part2_test() {
    assert_eq!(part2(TEST), "25272");
}
