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

#[inline_props]
pub(self) fn AnimateDice<'a>(
    cx: Scope<'a>,
    button_color: &'a str,
    button_text: &'a str,
    animate: &'a UseState<bool>,
) -> Element<'a> {
    render! {
        h1 {
            id: "rolling",
            class: "dice text-center",
            onanimationend: move |_| animate.set(false),
        }
        div {class: "grow"}
        button {
            class: "{button_color} rounded-md box-border w-full h-[15vh]",
            disabled: true,
            "{button_text}"
        }
    }
}
