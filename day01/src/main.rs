use std::fs;

fn main() {
    let lines = read_input();
    solution1(&lines);
    solution2(&lines);
}

fn solution1(lines: &Vec<u32>) {
    let mut current: u32 = 99_999;
    let mut count: u32 = 0;
    lines[..lines.len()-1]
        .iter()
        .for_each(|&n| {
            if n > current {
                count = count+1;
            }
            current = n;
        });

    println!("Solution 1: {}", count)
}

fn solution2(lines: &Vec<u32>) {
    let mut current: u32 = 99_999;
    let mut count: u32 = 0;

    for i in 0..lines.len()-2 {
        let pivot = lines[i] + lines[i+1] + lines[i+2];
        if pivot > current {
            count = count + 1;
        }
        current = pivot;
    }

    println!("Solution 2: {}", count)
}

fn read_input() -> Vec<u32> {
    let read = fs::read_to_string("input").unwrap();
    return read
        .split("\n")
        .map(|l| l.to_string())
        .map(|s| s.parse::<u32>())
        .filter(|p| p.is_ok())
        .map(|p| p.unwrap())
        .collect()
}
