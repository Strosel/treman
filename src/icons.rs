/// Icons from heroicons.com
use dioxus::prelude::*;

//BUG removing class from svg causes fill attribute to be skipped

pub fn QuestionMarkIcon(cx: Scope) -> Element {
    render! {
        svg {
            class: "icon",
            fill: "none",
            view_box: "0 0 24 24",
            stroke_width: "1.5",
            stroke: "#000",
            path {
                stroke_linecap: "round",
                stroke_linejoin: "round",
                d: "M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9 5.25h.008v.008H12v-.008z"
            }
        }
    }
}

pub fn LeftArrowIcon(cx: Scope) -> Element {
    render! {
        svg {
            class: "icon",
            fill: "none",
            view_box: "0 0 24 24",
            stroke_width: "1.5",
            stroke: "#000",
            path {
                stroke_linecap: "round",
                stroke_linejoin: "round",
                d: "M19.5 12h-15m0 0l6.75 6.75M4.5 12l6.75-6.75"
            }
        }
    }
}
