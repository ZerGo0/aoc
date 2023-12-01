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

    let output = lines
        .map(|line| {
            let mut chars = line.chars();
            // NOTE: find takes removes the element from the iterator
            let first = chars.find(|c| c.is_digit(10)).unwrap();
            let last = chars.rfind(|c| c.is_digit(10)).unwrap_or(first);

            // TODO: Is there a better way of doing this?
            let combined = format!("{}{}", first, last);
            let number = combined.parse::<usize>().unwrap();

            return number;
        })
        .sum();

    return output;
}

fn part_two(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();

    // one, two, three, four, five, six, seven, eight, nine
    let numbers_text = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    let output: usize = lines
        .map(|line| {
            let chars = line.chars();

            // The first and last digit can either be a number word or just a digit
            let mut first = '0';

            // find the first digit

            'first: for (i, c) in chars.clone().enumerate() {
                if c.is_digit(10) {
                    first = c;
                    break 'first;
                }

                if c == 'o' || c == 't' || c == 'f' || c == 's' || c == 'e' || c == 'n' {
                    let mut word = String::new();
                    word.push(c);

                    for c in chars.clone().skip(i + 1) {
                        if c.is_digit(10) {
                            first = c;
                            break 'first;
                        }

                        word.push(c);

                        if word.len() > 5 {
                            continue 'first;
                        }

                        if word.len() < 3 {
                            continue;
                        }

                        // Check if the word matches any of the numbersText elements
                        // if so return the index
                        let index = numbers_text.iter().position(|&x| x == word);
                        if index.is_some() {
                            // We need to add 1 to the index because we are skipping the first
                            first = std::char::from_digit((index.unwrap() + 1) as u32, 10).unwrap();
                            break 'first;
                        }
                    }
                }
            }

            let mut last = '0';

            'last: for (i, c) in line.chars().rev().enumerate() {
                if c.is_digit(10) {
                    last = c;
                    break 'last;
                }

                if c == 'e' || c == 'n' || c == 't' || c == 'r' || c == 'u' || c == 'o' {
                    let mut word = String::new();
                    word.push(c);

                    for c in line.chars().skip(line.len() - i) {
                        if c.is_digit(10) {
                            last = c;
                            break 'last;
                        }

                        word.push(c);
                        if word.len() > 5 {
                            continue 'last;
                        }

                        if word.len() < 3 {
                            continue;
                        }

                        // Check if the word matches any of the numbersText elements
                        // if so return the index
                        let index = numbers_text.iter().position(|&x| x == word);
                        if index.is_some() {
                            // We need to add 1 to the index because we are skipping the first
                            last = std::char::from_digit((index.unwrap() + 1) as u32, 10).unwrap();
                            break 'last;
                        }
                    }
                }
            }

            let combined = format!("{}{}", first, last);
            let number = combined.parse::<usize>().unwrap();

            //println!("{} -> {}", line, number);

            return number;
        })
        .sum();

    return output - 1;
}

#[allow(dead_code)]
const TEST_INPUT: &str = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

#[test]
fn test_part_one() {
    let out = part_one(Some(TEST_INPUT));

    assert_eq!(out, 142);
    println!("test_part_one: {}", out);
}

#[allow(dead_code)]
const TEST_INPUT2: &str = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

#[test]
fn test_part_two() {
    let out = part_two(Some(TEST_INPUT2));

    assert_eq!(out, 281);
    println!("test_part_two: {}", out)
}
