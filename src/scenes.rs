use dioxus::prelude::*;
use dioxus_router::prelude::*;

pub mod challange;
pub mod create;
pub mod game;
pub mod help;

pub use self::{challange::Challange, create::Create, game::Game, help::Help};

#[derive(Routable, Clone)]
pub enum Scene {
    #[route("/")]
    Game,
    #[route("/utmaning")]
    Challange,
    #[route("/ny-regel")]
    Create,
    #[route("/help")]
    Help,
}
