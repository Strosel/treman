use dioxus::prelude::*;

pub fn Challange(cx: Scope) -> Element {
    //TODO get dice state
    let d = (5, 6);
    render! {
        div {
            class: "flex justify-center items-center text-center",
            div {
                class: "flex flex-col gap-4 p-4 w-[100vmin] h-screen",
                h2 { class: "text-center", "Utmaning"}
                h1 {
                    class: "dice text-center",
                    span {class: "text-red-600", "{d.0}"}
                    span {class: "text-blue-600", "{d.1}"}
                }
                p { class: "grow", "Välj vars en tärning" }
                button {
                    class: "bg-secondary rounded-md box-border w-full h-[15vh]",
                    "Kör"
                }
            }
        }
    }
}
