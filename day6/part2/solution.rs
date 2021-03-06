use std::fs::File;
use std::io::{BufRead,BufReader};
use std::collections::HashMap;

fn main(){
    println!("Running Day 5 solution !");
    let file = File::open("./../input.txt").unwrap();
    let reader = BufReader::new(file);


    let mut map = HashMap::new();
    let mut total_questions = 0;
    let mut group_size = 0;
    for (_index, line) in reader.lines().enumerate(){

        let line_str = line.unwrap();
        if line_str.to_string() != ""{
            for c in line_str.chars() { 
                // do something with `c`
                if map.contains_key(&c) {
                    map.insert(c,map.get(&c).unwrap()+ &1);
                }else{
                    map.insert(c,1);
                }
            }
            group_size += 1;
        }else{
            for (_key , &value) in &map{
                if value == group_size {
                    total_questions += 1;
                }
            }

            group_size = 0;
            map.clear();
        }
    }

    println!("Total questions : {}", total_questions);
}