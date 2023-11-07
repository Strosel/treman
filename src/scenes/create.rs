use super::Scene;
use crate::components::*;
use crate::rules::*;
use dioxus::prelude::*;

pub fn Create(cx: Scope) -> Element {
    let trigger = use_state(cx, || RuleTrigger::Sum(2));
    let name = use_state(cx, String::new);

    render! {
        div {
            class: "flex flex-col gap-4 p-4 w-[100vmin] h-screen",
            Link{
                to: Scene::Game,
                class: "w-6 h-6",
                LeftArrowIcon{ }
            }

            fieldset {
                class: "flex gap-4 p-4 justify-center items-center",
                input {
                    r#type: "radio", name: "rule", id:"sum", checked: true,
                    oninput: move |_| trigger.modify(|v| match v {
                        RuleTrigger::Sum(_) => *v,
                        _ => RuleTrigger::Sum(2),
                    })
                }
                label {r#for: "sum", "Summa"}
                input {r#type: "radio", name: "rule", id: "pair",
                    oninput: move |_| trigger.modify(|v| match v {
                        RuleTrigger::Pair(..) => *v,
                        _ => RuleTrigger::Pair(2, 1),
                    })
                }
                label {r#for: "pair", "Par"}
                input {r#type: "radio", name: "rule", id: "single",
                    oninput: move |_| trigger.modify(|v| match v {
                        RuleTrigger::Single(_) => *v,
                        _ => RuleTrigger::Single(2),
                    })
                }
                label {r#for: "single", "En Tärning"}
            }

            match trigger.get() {
                RuleTrigger::Sum(v) => render!{
                    button {
                        class: "font-mono text-center text-black text-xl",
                        onclick: move |_| trigger.modify(|&v| v.inc(false)),
                        "{v}"
                    }
                },
                RuleTrigger::Single(v) => render!{
                    button {
                        class: "dice text-center text-black text-xl",
                        onclick: move |_| trigger.modify(|&v| v.inc(false)),
                        "{v}"
                    }
                },
                RuleTrigger::Pair(a, b) => render!{
                    div{
                        class: "flex justify-center items-center",
                        button {
                            class: "dice text-center text-black text-xl",
                            onclick: move |_| trigger.modify(|&v| v.inc(false)),
                            "{a}"
                        }
                        button {
                            class: "dice text-center text-black text-xl",
                            onclick: move |_| trigger.modify(|&v| v.inc(true)),
                            "{b}"
                        }
                    }
                },
                _ => unreachable!("Kan inte skapa ny treman regel"),
            }

            input {
                r#type: "text",
                placeholder:"Regel",
                value: name.get().as_str(),
                oninput: move |evt| name.set(evt.value.clone()),
            }
            div{ class: "grow" }
            button {
                class: "bg-secondary rounded-md box-border w-full h-[15vh]",
                disabled: name.get().is_empty(),
                onclick: move |_| {
                    if !name.get().is_empty() {
                        let rule = Rule::User {
                            trigger: *trigger.get(),
                            name: name.get().clone(),
                        };
                        use_shared_state::<Vec<Rule>>(cx)
                            .unwrap()
                            .write()
                            .push(rule);
                        *use_shared_state::<Scene>(cx)
                            .unwrap()
                            .write() = Scene::Game;
                    }
                },
                "Spara"
            }
        }
    }
}
