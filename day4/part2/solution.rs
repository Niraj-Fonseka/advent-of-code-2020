use std::fs::File;
use std::io::{BufRead,BufReader};
use std::collections::HashMap;

fn main(){
    println!("Running Day 4 solution !");


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
    
    let mut valid_count = 0;
    let mut count = 0;
    for (_index, line) in reader.lines().enumerate(){

        let line_str = line.unwrap();

        if line_str.to_string() == ""{

            if count == map.len() {
                valid_count +=1 
            }

            //lets reset
            count = 0;
            continue;
        }



        let line_split = line_str.split(" ").collect::<Vec<&str>>();

        for field in line_split.iter(){

                let key_value_split = field.split(":").collect::<Vec<&str>>();
                if map.contains_key(key_value_split[0]) {
          
                    if validate(key_value_split[0],&key_value_split[1]) {
                        count += 1;
                    }
                }
        }

    }

    println!("Total valid passports -> {}",valid_count);

}

fn validate(key: &str, value: &str) -> bool {
   
    if "byr" == key {

        let birth_year = value.parse::<i32>().unwrap();
        if birth_year >= 1920 && birth_year <= 2002 {
            return true;
        }
        return false;
        
    }else if "iyr" == key {

        let issued_year = value.parse::<i32>().unwrap();
        if issued_year >= 2010 && issued_year <= 2020 {
            return true;
        }
        return false;

    }else if "eyr" == key {

        let exp_year = value.parse::<i32>().unwrap();
        if exp_year >= 2020 && exp_year <= 2030 {
            return true;
        }
        return false;

    }else if "hgt" == key {

        if value.len() <= 2 {
            return false
        }

        let (metric_value,metric) = value.split_at(value.len()-2);
        let metric_value_num = metric_value.parse::<i32>().unwrap();

        if metric == "cm" {
            if metric_value_num >= 150 && metric_value_num <= 193 {
                return true;
            }
        }else if metric == "in"{
            if metric_value_num >= 59 && metric_value_num <= 76 {
                return true;
            }
        }
        return false;
    
    }else if "ecl" == key {

        if value == "amb" || value == "blu" || value == "brn"|| value == "gry"|| value == "grn"|| value == "hzl" || value == "oth"{
            return true;
        }
        return false;

    }else if "pid" == key {
        if value.len() == 9 && value.parse::<i32>().is_ok(){
            return true
        }

        return false;
    }else if "hcl" == key {
        if value.chars().next().unwrap() != '#'{
            return false
        }

        let color_split =  value.split("#").collect::<Vec<&str>>();

        if color_split[1].len() != 6 {
            return false;
        }
        let valid_color = i64::from_str_radix(color_split[1], 16);
        

        match valid_color {
            Ok(_valid_color)=> {
               return true;
            },
            Err(_e)=> {
                return false;
            }
        }
    }

    return true
}