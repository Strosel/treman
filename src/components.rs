use crate::{rules::*, scenes::Scene};
use dioxus::prelude::*;

mod fontawesome;
mod heroicons;
pub use self::{fontawesome::*, heroicons::*};

#[derive(Props)]
pub struct LinkProps<'a> {
    to: Scene,
    class: &'a str,
    children: Element<'a>,
}

pub fn Link<'a>(cx: Scope<'a, LinkProps<'a>>) -> Element<'a> {
    let LinkProps { to, class, .. } = cx.props;

    render! {
        a {
            class: *class,
            onclick: move |_| {
                *use_shared_state::<Scene>(cx).unwrap().write() = *to;
            },
            &cx.props.children
        }
    }
}

pub fn NavButton<'a>(cx: Scope<'a, LinkProps<'a>>) -> Element<'a> {
    let LinkProps { to, class, .. } = cx.props;

    render! {
        button{
            class: *class,
            onclick: move |_| {
                *use_shared_state::<Scene>(cx).unwrap().write() = *to;
            },
            &cx.props.children
        }
    }
}

#[inline_props]
pub fn DisplayTrigger<'a>(cx: Scope, trigger: &'a RuleTrigger) -> Element {
    match trigger {
        RuleTrigger::Sum(v) => render! {
            span {
                span {
                    class: "dice inline-block align-middle",
                    "0+0"
                }
                span {
                    class: "inline-block align-middle",
                    "={v}"
                }
            }
        },
        RuleTrigger::Pair(a, b) => render! {
            span {
                class: "dice inline-block align-middle",
                "{a}&{b}"
            }
        },
        RuleTrigger::Single(v) => render! {
            span {
                class: "dice inline-block align-middle",
                "{v}"
            }
        },
        RuleTrigger::Treman => render! {
            span {
                class: "dice inline-block align-middle",
                "3"
            }
        },
    }
}
