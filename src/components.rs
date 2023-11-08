use crate::scenes::Scene;
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
