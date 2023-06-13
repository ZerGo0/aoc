use std::str::Lines;

fn main() {
    let part = std::env::args().nth(1).expect("no part provided");

    let out = match part.as_str() {
        "1" => part_one(None),
        "2" => part_two(None),
        _ => panic!("invalid part provided"),
    };

    println!("Part {}: {}", part, out);
}

fn part_one(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();
    let line_count = lines.clone().count();

    let zero_bit_count: Vec<usize> = lines
        .clone()
        .map(|line| {
            line.chars()
                .enumerate()
                .filter(|(_, c)| *c == '0')
                .map(|(i, _)| i)
                .collect::<Vec<usize>>()
        })
        .flatten()
        // TODO: do we have to use a fixed sized vector here?
        .fold(vec![0; 12], |mut acc, i| {
            acc[i] += 1;
            acc
        });

    let (gamma_rate, epsilon_rate): (usize, usize) = (0..zero_bit_count.len())
        .map(|i| {
            if zero_bit_count[i] < line_count / 2 {
                (1, 0)
            } else {
                (0, 1)
            }
        })
        .fold((0, 0), |(gamma, epsilon), (g, e)| {
            (gamma * 2 + g, epsilon * 2 + e)
        });

    return gamma_rate * epsilon_rate;
}

fn part_two(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();

    let mut indexed_lines = get_indexed_lines(lines.clone().nth(0).unwrap().len());

    for (idx, line) in lines.clone().enumerate() {
        for (i, c) in line.chars().enumerate() {
            if c == '0' {
                indexed_lines[i][0].1.push(idx);
            } else {
                indexed_lines[i][1].1.push(idx);
            }
        }
    }

    let oxygen = get_index_common_bits(lines.clone(), &indexed_lines, true);
    let co2 = get_index_common_bits(lines.clone(), &indexed_lines, false);

    return oxygen * co2;
}

fn get_indexed_lines(size: usize) -> Vec<Vec<(usize, Vec<usize>)>> {
    let mut vec: Vec<Vec<(usize, Vec<usize>)>> = vec![];

    for _ in 0..size {
        vec.push(vec![(0, vec![]), (1, vec![])]);
    }

    return vec;
}

fn get_index_common_bits(
    lines: Lines,
    indexed_lines: &Vec<Vec<(usize, Vec<usize>)>>,
    most: bool,
) -> usize {
    let mut prev_results: Option<Vec<usize>> = None;
    let mut result: Vec<usize> = vec![];

    for (_, bits) in indexed_lines.iter().enumerate() {
        if prev_results.is_some() && prev_results.as_ref().unwrap().len() == 1 {
            return lines
                .clone()
                .nth(prev_results.unwrap()[0])
                .unwrap()
                .chars()
                .fold(0, |acc, c| if c == '0' { acc * 2 } else { acc * 2 + 1 });
        }

        let filtered_bits = bits
            .iter()
            .map(|(bit, values)| {
                (
                    bit,
                    values
                        .iter()
                        .filter(|v| {
                            prev_results.is_none() || prev_results.as_ref().unwrap().contains(v)
                        })
                        .map(|v| *v)
                        .collect::<Vec<usize>>(),
                )
            })
            .collect::<Vec<(&usize, Vec<usize>)>>();

        let bit: usize = if most {
            filtered_bits
                .iter()
                .max_by_key(|(_, values)| values.len())
                .unwrap()
                .0
                .clone()
        } else {
            filtered_bits
                .iter()
                .min_by_key(|(_, values)| values.len())
                .unwrap()
                .0
                .clone()
        };

        result.push(bit);

        prev_results = filtered_bits
            .iter()
            .filter(|(b, _)| **b == bit)
            .map(|(_, values)| values)
            .flatten()
            .map(|v| Some(*v))
            .collect::<Option<Vec<usize>>>();
    }

    return result.iter().fold(0, |acc, bit| acc * 2 + bit);
}

#[test]
fn test_part_one() {
    let test_input = "00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010";

    let out = part_one(Some(test_input));

    assert_eq!(out, 198);
    println!("test_part_one: {}", out);
}

#[test]
fn test_part_two() {
    let test_input = "00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010";

    let out = part_two(Some(test_input));

    assert_eq!(out, 230);
    println!("test_part_two: {}", out)
}
