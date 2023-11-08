/// Icons from heroicons.com
use dioxus::prelude::*;

//BUG removing class from svg causes attributes to be skipped

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

pub fn DownloadIcon(cx: Scope) -> Element {
    render! {
        svg {
            class: "icon",
            fill: "none",
            view_box: "0 0 24 24",
            stroke_width: "1.5",
            stroke: "#fff",
            path {
                stroke_linecap: "round",
                stroke_linejoin: "round",
                d: "M9 8.25H7.5a2.25 2.25 0 00-2.25 2.25v9a2.25 2.25 0 002.25 2.25h9a2.25 2.25 0 002.25-2.25v-9a2.25 2.25 0 00-2.25-2.25H15M9 12l3 3m0 0l3-3m-3 3V2.25"
            }
        }
    }
}

pub fn ShareIcon(cx: Scope) -> Element {
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
                d: "M9 8.25H7.5a2.25 2.25 0 00-2.25 2.25v9a2.25 2.25 0 002.25 2.25h9a2.25 2.25 0 002.25-2.25v-9a2.25 2.25 0 00-2.25-2.25H15m0-3l-3-3m0 0l-3 3m3-3V15"
            }
        }
    }
}

pub fn EllipsisH(cx: Scope) -> Element {
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
                d: "M6.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0zM12.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0zM18.75 12a.75.75 0 11-1.5 0 .75.75 0 011.5 0z"
            }
        }
    }
}

pub fn EllipsisV(cx: Scope) -> Element {
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
                d: "M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z"
            }
        }
    }
}
