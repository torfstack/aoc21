use std::fs;

fn main() {
    solution1();
    solution2();
}

fn solution2() {
    let input = read_input();
    let mut depth = 0;
    let mut pos = 0;
    let mut aim = 0;
    input.iter()
        .for_each(|i| {
            match i {
                Direction::FORWARD(n) => {
                    pos = pos + n;
                    depth = depth + (aim*n)
                },
                Direction::DOWN(n) => aim = aim + n,
                Direction::UP(n) => aim = aim - n
            }
        });
    println!("Solution 2 is {}", depth*pos)
}

fn solution1() {
    let input = read_input();
    let mut depth = 0;
    let mut pos = 0;
    input.iter()
        .for_each(|i| {
            match i {
                Direction::FORWARD(n) => pos = pos + n,
                Direction::DOWN(n) => depth = depth + n,
                Direction::UP(n) => depth = depth - n
            }
        });
    println!("Solution 1 is {}", depth*pos)
}

fn read_input() -> Vec<Direction> {
    let input = fs::read_to_string("input").unwrap();
    let lines: Vec<&str> = input.split("\n").collect();
    return lines[..lines.len()-1].into_iter()
        .map(|l| {
            let split: Vec<&str> = l.split(" ").collect();
            match split[0] {
                "forward" => Direction::FORWARD(split[1].parse::<usize>().unwrap()),
                "down" => Direction::DOWN(split[1].parse::<usize>().unwrap()),
                "up" => Direction::UP(split[1].parse::<usize>().unwrap()),
                &_ => panic!()
            }
        }).collect()
}

enum Direction {
    FORWARD(usize),
    DOWN(usize),
    UP(usize)
}
