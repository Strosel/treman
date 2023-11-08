use super::{AnimateDice, Scene};
use crate::{components::*, rules::*};
use dioxus::prelude::*;
use tinyrand::{RandRange, StdRand};

#[derive(Debug, Clone, PartialEq)]
enum OverlayState {
    Install,
    Rule(Rule),
    Hidden,
}

#[inline_props]
fn DisplayRule(cx: Scope, rule: Rule) -> Element {
    match rule {
        Rule::Base {
            trigger,
            name,
            desc,
        } => render! {
            p {
                class: "text-sm",
                DisplayTrigger { trigger: trigger }
            }
            p {
                class: "text-sm",
                "{name}"
            }
            p {
                class: "text-xs",
                "{desc}"
            }

        },
        Rule::User { trigger, name } => render! {
            p {
                class: "text-sm",
                DisplayTrigger { trigger: trigger }
            }
            p {
                class: "text-sm",
                "{name}"
            }
            p {
                class: "text-xs",
                "Detta är en egen regel. Fråga den som skapat den."
            }
        },
    }
}

#[inline_props]
fn PlayGame<'a>(
    cx: Scope<'a>,
    dice: &'a UseState<(u8, u8)>,
    animate: &'a UseState<bool>,
    overlay: &'a UseState<OverlayState>,
) -> Element<'a> {
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
            for r in rules.iter().cloned() {
                li {
                    class: "cursor-pointer",
                    onclick: move |_| overlay.set(OverlayState::Rule(r.clone())),
                    r.name()
                }
            }
        }

        if **dice == (6,6) {
            render! {
                NavButton {
                    class: "bg-secondary",
                    to: Scene::Create,
                    "Jag har aldrig sett..."
                }
            }
        }

        if **dice == (1,2) || **dice == (2,1) {
            render! {
                NavButton {
                    class: "bg-secondary",
                    to: Scene::Challange,
                    "Utmaning"
                }
            }
        } else {
            render! {
                button {
                    class: "bg-primary",
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
    let overlay = use_state(cx, || OverlayState::Hidden);

    render! {
        if **overlay != OverlayState::Hidden {
            render! {
                div {
                    class: "overlay bg-neutral-100 rounded-lg shadow-2xl p-4 min-h-[40vh] text-left flex flex-col justify-center",
                    match **overlay {
                        OverlayState::Install => render!{
                            p {
                                class: "text-sm mb-2",
                                "Installera Treman 📲"
                            }
                            p {
                                class: "text-xs",
                                "Om du vill spela Treman även offline kan du installera appen till din telefon."
                            }
                            ol {
                                class: "list-inside list-decimal my-4 text-xs",
                                li {
                                    class: "w-fit",
                                    "Välj 'dela' (ser ut som ikonerna här)"
                                    div {
                                        class : "flex gap-6 justify-evenly",
                                        span { class: "icon", ShareIcon { } }
                                        span { class: "icon", EllipsisH { } }
                                        span { class: "icon", EllipsisV { } }
                                    }
                                }
                                li { "Välj 'Lägg till på hemskärm' eller 'Installera'." }
                                li { "Njut av Treman offline 🎲🍺" }
                            }
                        },
                        OverlayState::Rule(ref r) => render!{
                            DisplayRule { rule: r.clone() }
                        },
                        OverlayState::Hidden => unreachable!(),
                    }
                }
            }
        }

        div {
            class: "z-0 flex flex-col gap-4 p-4 w-[100vmin] h-screen",
            onclick: move |_| if **overlay != OverlayState::Hidden {
                overlay.set(OverlayState::Hidden)
            },
            div {
                class: "flex flex-row-reverse h-[4vh] w-100 justify-between items-center",

                Link{
                    to: Scene::Help,
                    class: "icon",
                    QuestionMarkIcon { }
                }

                button {
                    id: "install",
                    class: "bg-secondary pl-3 pr-2 w-fit h-fit min-h-fit text-sm inline-flex items-center gap-2",
                    onclick: move |_| overlay.set(OverlayState::Install),
                    "Installera"
                    span{
                        class: "icon",
                        DownloadIcon { }
                    }
                }

                span { /* empty span makes the install/update button centered */ }
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
                        overlay: overlay,
                    }
                }
            }
        }
    }
}
