use super::{AnimateDice, Scene};
use crate::icons::*;
use crate::rules::*;
use dioxus::prelude::*;
use dioxus_router::prelude::*;
use tinyrand::{RandRange, StdRand};

#[inline_props]
fn PlayGame<'a>(
    cx: Scope<'a>,
    dice: &'a UseState<(u8, u8)>,
    animate: &'a UseState<bool>,
) -> Element<'a> {
    let nav = use_navigator(cx);
    let rules: Vec<_> = use_shared_state::<Vec<Rule>>(cx)
        .unwrap()
        .read()
        .iter()
        .filter(|r| r.check(dice))
        .cloned()
        .collect();

    render! {
        h1 {
            class: "dice text-center",
            "{dice.0}{dice.1}"
        }
        ul {
            class: "grow scrollbox",
            if rules.is_empty() {
                render! {
                    li { "Ingenting" }
                }
            }
            for r in rules.iter() {
                li { r.name() }
            }
        }

        if **dice == (6,6) {
            render! {
                button {
                    class: "bg-secondary rounded-md box-border w-full h-[15vh] min-h-[15vh]",
                    onclick: move |_| { nav.replace(Scene::Create); },
                    "Jag har aldrig sett..."
                }
            }
        }

        if **dice == (1,2) || **dice == (2,1) {
            render! {
                button {
                    class: "bg-secondary rounded-md box-border w-full h-[15vh] min-h-[15vh]",
                    onclick: move |_| { nav.replace(Scene::Challange); },
                    "Utmaning"
                }
            }
        } else {
            render! {
                button {
                    class: "bg-primary rounded-md box-border w-full h-[15vh] min-h-[15vh]",
                    onclick: move |_| {
                        let mut rng = use_shared_state::<StdRand>(cx).unwrap().write_silent();
                        dice.set((
                            rng.next_range(1..7_u16) as u8,
                            rng.next_range(1..7_u16) as u8
                        ));
                        animate.set(true);
                    },
                    "Rulla"
                }
            }
        }
    }
}

pub fn Game(cx: Scope) -> Element {
    let dice = use_state(cx, || (0, 0));
    let animate = use_state(cx, || false);

    render! {
        div {
            class: "flex flex-col gap-4 p-4 w-[100vmin] h-screen",
            Link{
                to: Scene::Help,
                class: "w-6 h-6 self-end",
                QuestionMarkIcon { }
            }

            if **animate {
                render!{
                    AnimateDice {
                        animate: animate,
                        button_color: "bg-primary",
                        button_text: "Rulla",
                    }
                }
            } else {
                render!{
                    PlayGame {
                        dice: dice,
                        animate: animate,
                    }
                }
            }
        }
    }
}
