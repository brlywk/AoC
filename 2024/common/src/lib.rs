use anyhow::*;
use std::{fs::read_to_string, path::Path};

#[derive(Debug)]
pub struct Input {
    pub content: String,
    pub lines: Vec<String>,
}

impl Input {
    pub fn new(input_path: impl AsRef<Path>) -> Self {
        let content = read_file(&input_path).unwrap_or_else(|err| {
            panic!(
                "Failed to open file '{}': {}",
                input_path.as_ref().display(),
                err
            )
        });

        return Self {
            lines: content.lines().map(String::from).collect(),
            content,
        };
    }
}

fn read_file(input_path: impl AsRef<Path>) -> Result<String> {
    Ok(read_to_string(input_path)?)
}

pub mod prelude {
    pub use super::Input;
    pub use anyhow;
}
