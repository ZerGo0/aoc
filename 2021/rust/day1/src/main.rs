#![feature(array_windows)]

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

    let increase = lines
        .map(|line| line.parse::<usize>().unwrap())
        .fold((0, 0), |(prev, count), next| {
            if prev == 0 {
                (next, count)
            } else if next >= prev {
                (next, count + 1)
            } else {
                (next, count)
            }
        })
        .1;

    return increase;
}

fn part_two(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input
        .lines()
        .map(|line| line.parse::<usize>().unwrap())
        .collect::<Vec<usize>>();

    let increase = lines
        .array_windows::<3>()
        .map(|[a, b, c]| (a + b + c))
        .fold((0, 0), |(prev, count), curr| -> (usize, usize) {
            match prev {
                0 => (curr, count),
                _ => {
                    if curr > prev {
                        (curr, count + 1)
                    } else {
                        (curr, count)
                    }
                }
            }
        })
        .1;

    return increase;
}

#[test]
fn test_part_one() {
    let test_input = "199
200
208
210
200
207
240
269
260
263";

    let out = part_one(Some(test_input));

    assert_eq!(out, 7);
    println!("test_part_one: {}", out);
}

#[test]
fn test_part_two() {
    let test_input = "199
200
208
210
200
207
240
269
260
263";

    let out = part_two(Some(test_input));

    assert_eq!(out, 5);
    println!("test_part_two: {}", out)
}
