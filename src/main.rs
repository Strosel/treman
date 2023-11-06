use dioxus::prelude::*;
use crate::rules::*;

mod rules;

fn main() {
    // init debug tool for WebAssembly
    wasm_logger::init(wasm_logger::Config::default());
    console_error_panic_hook::set_once();

    dioxus_web::launch(app);
}

fn app(cx: Scope) -> Element {
    render!{
        Help{}
    }
}

fn Help(cx: Scope) -> Element {
    render! {
        h2{
            class: "p-4",
            "Regler"
        }
        p {
            class: "text-xs px-4 py-2",
            "Treman är ett dryckesspel som går ut på att slå tärningar för att bestämma vem som dricker och hur mycket."
        }
        h3 { 
            class: "p-4",
            "Vem är treman?"
        }
        p {
            class: "text-xs px-4 py-2",
            "Börja med att utse en spelare till 'treman', vem som är treman kommer att ändras under spelets gång. Bara en person kan vara treman i taget."
        }
        p {
            class: "text-xs p-4 py-2",
            "Följande leder till att en ny spelare blir 'treman':"
        }
        ul {
            class: "list-disc text-xs px-4 py-2 mx-4",
            li {"Om en ny person går med i spelet är hen nu treman."}
            li {"Om en person lämnar bordet och kommer tillbaka (t.ex. går på toa eller hämtar mer dricka) är hen nu treman."}
            li {"Skulle treman lämna bordet, blir föregående person treman igen."}
            li {"Vissa tärningsslag kan också resultera i en ny treman."}
        }
        p {
            class: "text-xs p-4 py-2",
            "När en ny person blir treman skålar nya och gamla treman och tar vars en klunk."
        }
        h3 { 
            class: "p-4",
            "Sen då?"
        }
        p {
            class: "text-xs p-4 py-2",
            "När ni valt treman kan spelet börja. En spelare slår tärningarna och ropar ut de regler som de slagit. En regel kan påverka en eller flera personer, se listan nedan, när alla påverkade har slutfört regeln slås tärningarna på nytt. Samma spelare fortsätter alltid att slå tärningarna tills dess att de slår 'Ingenting', då skickas tärningarna vidare medsols."
        }
        h3 { 
            class: "p-4",
            "Regellista"
        }
        //TODO
        p {
            class: "dice text-xs p-4 py-2",
            "3+2"
        }
    }
}

fn Create(cx: Scope) -> Element {
    let ty = use_state(cx, || RuleTrigger::Sum(2));

    render! {
        div {
            class: "flex justify-center items-center text-center",
            div {
                class: "flex flex-col gap-4 p-4 w-[100vmin] h-screen",

                fieldset {
                    class: "flex gap-4 p-4 justify-center items-center",
                    input {
                        r#type: "radio", name: "rule", id:"sum", checked: true,
                        oninput: move |_| ty.modify(|v| match v {
                            RuleTrigger::Sum(_) => *v,
                            _ => RuleTrigger::Sum(2),
                        })
                    }
                    label {r#for: "sum", "Summa"}
                    input {r#type: "radio", name: "rule", id: "pair", 
                        oninput: move |_| ty.modify(|v| match v {
                            RuleTrigger::Pair(..) => *v,
                            _ => RuleTrigger::Pair(2, 1),
                        })
                    }
                    label {r#for: "pair", "Par"}
                    input {r#type: "radio", name: "rule", id: "single", 
                        oninput: move |_| ty.modify(|v| match v {
                            RuleTrigger::Single(_) => *v,
                            _ => RuleTrigger::Single(2),
                        })
                    }
                    label {r#for: "single", "En Tärning"}
                }

                match ty.get() {
                    RuleTrigger::Sum(v) => render!{
                        button {
                            class: "font-mono text-center text-black text-xl", 
                            onclick: move |_| ty.modify(|&v| v.inc(false)),
                            "{v}"
                        }
                    },
                    RuleTrigger::Single(v) => render!{
                        button {
                            class: "dice text-center text-black text-xl", 
                            onclick: move |_| ty.modify(|&v| v.inc(false)),
                            "{v}"
                        }
                    },
                    RuleTrigger::Pair(a, b) => render!{
                        div{
                            class: "flex justify-center items-center",
                            button {
                                class: "dice text-center text-black text-xl", 
                                onclick: move |_| ty.modify(|&v| v.inc(false)),
                                "{a}"
                            }
                            button {
                                class: "dice text-center text-black text-xl", 
                                onclick: move |_| ty.modify(|&v| v.inc(true)),
                                "{b}"
                            }
                        }
                    }
                }

                input { r#type: "text", placeholder:"Regel" }
                div{ class: "grow" }
                button { 
                    class: "bg-secondary rounded-md box-border w-full h-[15vh]",
                    "Spara" 
                }
            }
        }
    }
}

fn Challange(cx: Scope) -> Element {
    //TODO get dice state
    let d = (5,6);
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

fn Game(cx: Scope) -> Element {
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
