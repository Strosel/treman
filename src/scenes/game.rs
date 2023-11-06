use crate::rules::*;
use dioxus::prelude::*;
use tinyrand::{RandRange, StdRand};

pub fn Game(cx: Scope) -> Element {
    let d = use_state(cx, || (0, 0));
    let rules: Vec<_> = use_shared_state::<Vec<Rule>>(cx)
        .unwrap()
        .read()
        .iter()
        .filter(|r| r.check(d))
        .cloned()
        .collect();

    render! {
        div {
            class: "flex justify-center items-center text-center",
            div {
                class: "flex flex-col gap-4 p-4 w-[100vmin] h-screen",
                h1 { class: "dice text-center", "{d.0}{d.1}"}
                ul {
                    class: "grow",
                    if rules.is_empty() {
                        render! {
                            li { "Ingenting" }
                        }
                    }
                    for r in rules.iter() {
                        render! {
                            li { r.name() }
                        }
                    }
                }
                if *d == (6,6) {
                    render! {
                        button {
                            class: "bg-secondary rounded-md box-border w-full h-[15vh]",
                            "Jag har aldrig sett..."
                        }
                    }
                }

                if *d == (1,2) || *d == (2,1) {
                    render! {
                        button {
                            class: "bg-secondary rounded-md box-border w-full h-[15vh]",
                            "Utmaning"
                        }
                    }
                } else {
                    render! {
                        button {
                            class: "bg-primary rounded-md box-border w-full h-[15vh]",
                            onclick: move |_| {
                                let mut rng = use_shared_state::<StdRand>(cx).unwrap().write_silent();
                                d.set((
                                    rng.next_range(1..7_u16) as u8,
                                    rng.next_range(1..7_u16) as u8
                                ));
                            },
                            "Rulla"
                        }
                    }
                }
            }
        }
    }
}
