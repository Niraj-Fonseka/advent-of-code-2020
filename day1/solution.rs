use std::fs::File;
use std::io::{BufRead, BufReader};

fn main(){
    println!("Running day1 solution !");


    let file  = File::open("input2.txt").unwrap();

    let reader = BufReader::new(file);

    let mut number_lst = vec![1i32];

    for (_index , line) in reader.lines().enumerate()  {
        let line_str = line.unwrap();

        let to_number = line_str.parse::<i32>().unwrap();
        number_lst.push(to_number);
    }


    for i in 0..number_lst.len(){
        let step = i + 1;
        for x in step..number_lst.len(){
            let total = number_lst[i] + number_lst[x];

            if total == 2020 {
                println!("{} and {}", number_lst[i] , number_lst[x]);
                println!("multiplied : {} ", number_lst[i] * number_lst[x]);
                break;
            }
        }
    }


}