use std::fs::File;
use std::io::{BufRead, BufReader};

fn main(){
    println!("Running day1 solution !");
    
    let file  = File::open("./../input.txt").unwrap();
    let reader = BufReader::new(file);
    let mut number_lst = vec![1i32];

    for (_index , line) in reader.lines().enumerate()  {
        let line_str = line.unwrap();
        let to_number = line_str.parse::<i32>().unwrap();
        number_lst.push(to_number);
    }

    //O(n^3)
    unoptimized(number_lst);

}

fn unoptimized(number_lst: Vec<i32>){
    for i in 0..number_lst.len(){
        let step = i + 1;
        for x in step..number_lst.len(){
            let step_two = x + 1;

            for y in step_two..number_lst.len(){
                let total = number_lst[i] + number_lst[x] + number_lst[y];
                if total == 2020 {
                    println!("{} and {} and {} ", number_lst[i] , number_lst[x], number_lst[y]);
                    println!("multiplied : {} ", number_lst[i] * number_lst[x]* number_lst[y]);
                    break;
                }
            }
        }
    }
}