use std::fs::File;
use std::io::{BufRead,BufReader};
use std::collections::HashMap;

fn main(){
    println!("Running Day 4  solution !");


    let file = File::open("./../input.txt").unwrap();
    let reader = BufReader::new(file);
    let mut map = HashMap::new();
    map.insert("byr", true);
    map.insert("iyr", true);
    map.insert("eyr", true);
    map.insert("hgt", true);
    map.insert("hcl", true);
    map.insert("ecl", true);
    map.insert("pid", true);
    map.insert("cid", true);
    
    let mut valid_count = 0;
    let mut cid = false;
    let mut count = 0;
    for (_index, line) in reader.lines().enumerate(){

        let line_str = line.unwrap();

        if line_str.to_string() == ""{
            if count == map.len() {
                valid_count +=1 
            }else if (count == map.len()-1) && cid == false {
                valid_count +=1 
            }

            //lets reset
            cid = false;
            count = 0;
            continue;
        }



        let line_split = line_str.split(" ").collect::<Vec<&str>>();

        for field in line_split.iter(){

                let key_value_split = field.split(":").collect::<Vec<&str>>();
                if map.contains_key(key_value_split[0]) {
                    if key_value_split[0] == "cid" {
                        cid = true;
                    }
                    count += 1;
                }
        }

    }

    println!("Total valid passports -> {}",valid_count);

}