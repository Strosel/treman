pub mod challange;
pub mod create;
pub mod game;
pub mod help;

pub use self::{challange::Challange, create::Create, game::Game, help::Help};

pub enum Scene {
    Game,
    Challange,
    Create,
    Help,
}
