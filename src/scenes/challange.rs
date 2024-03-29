use super::{AnimateDice, Scene};
use crate::components::*;
use dioxus::prelude::*;
use std::cmp::Ordering;
use tinyrand::{RandRange, StdRand};

#[inline_props]
fn PlayChallange<'a>(
    cx: Scope<'a>,
    dice: &'a UseState<(u8, u8)>,
    animate: &'a UseState<bool>,
) -> Element<'a> {
    render! {
        h1 {
            class: "dice text-center",
            span {class: "text-red-600", "{dice.0}"}
            span {class: "text-blue-600", "{dice.1}"}
        }
        match Ord::cmp(&dice.0, &dice.1) {
            Ordering::Equal => render!{
                p {
                    class: "grow",
                    "Välj vars en tärning, den utmanade väljer först"
                }
            },
            Ordering::Less => render!{
                p { class: "grow", "Blå är ny Treman" }
            },
            Ordering::Greater => render!{
                p { class: "grow", "Röd är ny Treman" }
            },
        }

        if dice.0 == dice.1 {
            render!{
                button {
                    class: "bg-secondary",
                    onclick: move |_| {
                        let mut rng = use_shared_state::<StdRand>(cx).unwrap().write_silent();
                        let (mut red, mut blue) = (0, 0);
                        while red == blue {
                            red = rng.next_range(1..7_u16) as u8;
                            blue = rng.next_range(1..7_u16) as u8;
                        }
                        dice.set((red, blue));
                        animate.set(true);
                    },
                    "Kör"
                }
            }
        } else {
            render!{
                NavButton {
                    class: "bg-secondary",
                    to: Scene::Game,
                    "Ok"
                }
            }
        }
    }
}

pub fn Challange(cx: Scope) -> Element {
    let dice = use_state(cx, || (0, 0));
    let animate = use_state(cx, || false);

    render! {
        div {
            class: "flex flex-col gap-4 p-4 w-[100vmin] h-screen",
            Link {
                to: Scene::Game,
                class: "icon",
                LeftArrowIcon{ }
            }

            if **animate {
                render!{
                    AnimateDice {
                        animate: animate,
                        button_color: "bg-secondary",
                        button_text: "Kör",
                    }
                }
            } else {
                render!{
                    PlayChallange {
                        dice: dice,
                        animate: animate,
                    }
                }
            }
        }
    }
}
