use dioxus::prelude::*;

fn main() {
    // init debug tool for WebAssembly
    wasm_logger::init(wasm_logger::Config::default());
    console_error_panic_hook::set_once();

    dioxus_web::launch(app);
}

fn app(cx: Scope) -> Element {
    //TODO get dice state
    let d = (5,6);
    render! {
        div {
            class: "flex justify-center items-center text-center",
            div {
                class: "flex flex-col gap-4 p-4 w-[100vmin] h-screen",
                h1 { class: "dice text-center", "{d.0}{d.1}"}
                p { class: "grow", "rules" }
                button { 
                    class: "bg-secondary rounded-md box-border w-full h-[15vh]",
                    "Jag har aldrig sett..." 
                }
                button { 
                    class: "bg-primary rounded-md box-border w-full h-[15vh]",
                    "Rulla" 
                }
            }
        }
    }
}
