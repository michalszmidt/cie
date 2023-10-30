use clap::{Arg, ArgAction, Command};

// COMMANDS
pub fn get_command_modify() -> Command {
    let command_modify: Command = Command::new("modify")
        .short_flag('M')
        .long_flag("modify")
        .about("Modify ICal");
    return command_modify;
}

// ARGS
pub fn get_args_modify() -> Vec<Arg> {
    let arg_path: Arg = Arg::new("path")
        .help("Path to file to be read [path without quotes]")
        .short('p')
        .long("path")
        .action(ArgAction::Set);

    let arg_config: Arg = Arg::new("config")
        .help("Path to config file [path without quotes]")
        .short('c')
        .long("config")
        .action(ArgAction::Set);

    let arg_out: Arg = Arg::new("out")
        .help("Path to out file [path without quotes]")
        .short('o')
        .long("out")
        .action(ArgAction::Set);

    return vec![arg_path, arg_config, arg_out];
}
