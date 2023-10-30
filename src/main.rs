pub mod commands;
pub mod processing;
pub mod structs;

use clap::{parser::ValuesRef, Command};
use commands::{get_args_modify, get_command_modify};
use processing::process;

const VERSION: &str = env!("CARGO_PKG_VERSION");

fn main() {
    let cie = Command::new("cie")
        .about("CLI ICal editor")
        .version(VERSION)
        .subcommand_required(true)
        .arg_required_else_help(true)
        .author("Michał Szmidt")
        .subcommand(get_command_modify().args(get_args_modify()))
        .get_matches();

    if let Some(("modify", query_matches)) = cie.subcommand() {
        let mut out = "./out.txt".to_string();
        let mut path = "NAN".to_string();
        let mut config = "./cie.yaml".to_string();

        if let Some(value_of_out) = query_matches.get_many::<String>("out") {
            out = get_param(value_of_out);
        }
        if let Some(value_of_path) = query_matches.get_many::<String>("path") {
            path = get_param(value_of_path);
        }

        if let Some(value_of_config) = query_matches.get_many::<String>("config") {
            config = get_param(value_of_config);
        }

        let count_removed = process(path, config, out);
        println!("Removed Items: {}", count_removed);
    } else {
        unreachable!()
    }
}

fn get_param(valuesref: ValuesRef<String>) -> String {
    let x = valuesref.map(|s| s.as_str()).collect::<Vec<_>>().join(", ");
    return x;
}
