use std::fs::File;
use std::io::{BufRead, BufReader};
//use std::any::type_name;

fn main(){
    println!("Running day2 solution !");

    let file = File::open("./../input.txt").unwrap();
    let reader = BufReader::new(file);

    let mut total = 0;
    for (_index , line) in reader.lines().enumerate()  {
        let line_str = line.unwrap();

        let number_str_vec = line_str.split(":").collect::<Vec<&str>>();

        let char_to_range_vec = number_str_vec[0].split(" ").collect::<Vec<&str>>();

        let split_min_max_range_vec = char_to_range_vec[0].split("-").collect::<Vec<&str>>();

        let min = split_min_max_range_vec[0].parse::<i32>().unwrap();
        let max = split_min_max_range_vec[1].parse::<i32>().unwrap();

        let mut store:[i32;256] = [0;256];

        for c in number_str_vec[1].trim().chars(){
            let index = c as usize - 48;
            store[index] = store[index] + 1;
        }

        let c_check = char_to_range_vec[1].chars().nth(0).unwrap();
        let c_check_to_index = c_check as usize - 48;


        let mentions = store[c_check_to_index];

        if mentions >= min && mentions <= max {
            println!("Match !");
            println!("str : {}", number_str_vec[1].trim());
            println!("char : {}", char_to_range_vec[1]);
            println!("min : {}", min);
            println!("max : {}",max);
            println!();
            total = total + 1;
        }
    }

    println!("Total matches : {} ", total);

}
