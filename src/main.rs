use crate::{rules::*, scenes::*};
use dioxus::prelude::*;
use dioxus_router::prelude::*;
use tinyrand::{Seeded, StdRand};

mod icons;
mod rules;
mod scenes;

fn main() {
    // init debug tool for WebAssembly
    wasm_logger::init(wasm_logger::Config::default());
    console_error_panic_hook::set_once();

    dioxus_web::launch(app);
}

fn app(cx: Scope) -> Element {
    //Use js_sys::Math::random to seed tinyrand which has better api
    use_shared_state_provider(cx, || {
        StdRand::seed((js_sys::Math::random() * u16::MAX as f64) as u64)
    });
    use_shared_state_provider(cx, || Rule::BASE.to_vec());

    render! {
        Router::<Scene> {}
    }
}
