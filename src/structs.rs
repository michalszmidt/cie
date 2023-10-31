use serde::Deserialize;

#[derive(Clone, Deserialize)]
pub struct CIE {
    pub settings: Settings,
    pub to_remove: Vec<REMOVE>,
    pub to_move: Vec<MOVE>,
    // pub to_modify: Vec<MODIFY>,
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

#[derive(Clone, Deserialize)]
pub struct MOVE {
    pub regex: String,
    pub key: String,
    pub duration: String, // https://github.com/waltzofpearls/dateparser/tree/main#accepted-date-formats
    pub duration_unit: String,
    pub at_weekday: String,
}

// #[derive(Clone, Deserialize)]
// pub struct MODIFY {
//     pub regex: String,
//     pub key: String,
//     pub action: String,
//     pub action_key: String,
// }
