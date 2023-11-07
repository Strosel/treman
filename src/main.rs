#![allow(non_snake_case)]

use crate::{rules::*, scenes::*};
use dioxus::prelude::*;
use tinyrand::{Seeded, StdRand};

mod components;
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
    use_shared_state_provider(cx, Scene::default);

    match *use_shared_state::<Scene>(cx).unwrap().read() {
        Scene::Game => render! { Game { } },
        Scene::Challange => render! { Challange { } },
        Scene::Create => render! { Create { } },
        Scene::Help => render! { Help{ } },
    }
}
