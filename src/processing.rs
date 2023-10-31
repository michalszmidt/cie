use std::borrow::BorrowMut;
use std::collections::HashSet;
use std::fs::{read_to_string, File};
use std::io::{self, BufWriter, Write};
use std::str::FromStr;

use chrono::{DateTime, Datelike, Duration};
// use chrono::{self, DateTime, TimeZone};
// use chrono_tz::{self, UTC};
// use dateparser::*;

use icalendar::{Calendar, Component, EventLike};
use rayon::prelude::{IntoParallelRefIterator, IntoParallelRefMutIterator, ParallelIterator};
use regex::Regex;

use crate::structs::CIE;

pub fn process(path: String, config: String, out: String) -> usize {
    let calendar_as_string = read_to_string(path).unwrap();
    let mut counter: usize = 0;
    let file_rejected = file_write(out).unwrap();
    let mut writer_rejected = BufWriter::new(file_rejected);

    let mut ical = Calendar::from_str(calendar_as_string.as_str()).unwrap();

    let cie_yaml_exact: Option<CIE> = match serde_yaml::from_reader(file_to_lines(config).unwrap())
    {
        Ok(x) => x,
        Err(e) => {
            println!("{}", e);
            None
        }
    };

    let cie_yaml = &cie_yaml_exact.unwrap();

    let to_remove = cie_yaml.clone().to_remove;
    let to_move = cie_yaml.clone().to_move;

    ical.components.retain(|x| match x.as_event() {
        Some(event) => {
            let set: HashSet<_> = to_remove
                .par_iter()
                .filter(|remove_rule| event.property_value(&remove_rule.key).is_some())
                .map(|remove_rule| {
                    let rule = Regex::new(&remove_rule.regex).unwrap();
                    let value = event.property_value(&remove_rule.key).unwrap();
                    return rule.is_match(value);
                })
                .collect();
            if set.contains(&true) {
                counter += 1;
                return false;
            }
            return true;
        }
        None => true,
    });

    //Must operate on events clones
    // Below wont work.
    ical.components
        .iter_mut()
        .filter_map(|x| x.as_event())
        .for_each(|event| {
            to_move.iter().for_each(|remove_rule| {
                let time: DateTime<_> = event.clone().get_timestamp().unwrap();

                let rule = Regex::new(&remove_rule.regex).unwrap();
                let value = event
                    .clone()
                    .property_value(&remove_rule.key)
                    .unwrap()
                    .to_string();

                if rule.is_match(value.as_str()) {
                    if remove_rule.at_weekday.as_str() == ""
                        || remove_rule.at_weekday.as_str()
                            == time.weekday().to_string().to_lowercase().as_str()
                    {
                        let duration = match remove_rule.duration_unit.as_str() {
                            "h" => Duration::hours(remove_rule.duration.parse().unwrap()),
                            "m" => Duration::minutes(remove_rule.duration.parse().unwrap()),
                            _ => Duration::hours(remove_rule.duration.parse().unwrap()),
                        };

                        time.checked_add_signed(duration);
                        event.starts(time);
                    }
                }
            })
        });

    ical = ical.done();

    let to_save: String = ical.to_string();
    _ = writer_rejected.write_all(to_save.as_bytes());
    _ = writer_rejected.flush();

    return counter;
}

pub fn file_to_lines(path: String) -> io::Result<File> {
    let file = File::open(path)?;
    return Ok(file);
}

pub fn file_write(path: String) -> io::Result<File> {
    let file = File::create(path)?;
    return Ok(file);
}
