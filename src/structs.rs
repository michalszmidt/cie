use serde::Deserialize;

#[derive(Clone, Deserialize)]
pub struct CIE {
    pub settings: Settings,
    pub to_remove: Vec<REMOVE>,
    // pub to_move: Vec<MOVE>,
}

#[derive(Clone, Deserialize)]
pub struct Settings {
    pub timezone: String,
}

#[derive(Clone, Deserialize)]
pub struct REMOVE {
    pub regex: String,
    pub key: String,
}

// #[derive(Clone, Deserialize)]
// pub struct MOVE {
//     pub regex: String,
//     pub key: String,
//     pub start_time: String, // hr:min
//     pub end_time: String,   // hr:min
// }
