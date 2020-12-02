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

        let min = split_min_max_range_vec[0].parse::<usize>().unwrap()-1;
        let max = split_min_max_range_vec[1].parse::<usize>().unwrap()-1;

        let c_check = char_to_range_vec[1].chars().nth(0).unwrap();
        // let c_check_to_index = c_check as usize - 48;
        let password = number_str_vec[1].trim();

        
        let min_char = password.chars().nth(min).unwrap();
        let max_char = password.chars().nth(max).unwrap();


        if (min_char == c_check && max_char != c_check) || (min_char != c_check && max_char == c_check){
            println!("Match !!");
            println!("Min : {} , Max : {}",min,max);
            println!("char to check : {}", c_check);
            println!("Min char : {} , Max char : {}\n", min_char , max_char);
            total = total  + 1;
        }
    }

    println!("Total Valid : {}", total);

}

