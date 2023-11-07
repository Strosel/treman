use super::Scene;
use crate::components::*;
use crate::rules::*;
use dioxus::prelude::*;

#[inline_props]
fn DisplayTrigger<'a>(cx: Scope, trigger: &'a RuleTrigger) -> Element {
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

#[inline_props]
fn DisplayRule(cx: Scope, rule: Rule) -> Element {
    match rule {
        Rule::Base {
            trigger,
            name,
            desc,
        } => render! {
            div {
                p {
                    class: "text-sm",
                    DisplayTrigger { trigger: trigger }
                    span {
                        class: "inline-block align-middle px-4",
                        "{name}"
                    }
                }
                p {class: "text-xs", "{desc}"}
            }
        },
        Rule::User { trigger, name } => render! {
            p {
                class: "text-sm",
                DisplayTrigger { trigger: trigger }
                span {
                    class: "inline-block align-middle px-4",
                    "{name}"
                }
            }
        },
    }
}

pub fn Help(cx: Scope) -> Element {
    let rules: Vec<_> = use_shared_state::<Vec<Rule>>(cx)
        .unwrap()
        .read()
        .iter()
        .cloned()
        .collect();

    render! {
        div {
            class: "flex flex-col text-left gap-4 p-4 w-[100vmin] h-screen",
            Link{
                to: Scene::Game,
                class: "w-6 h-6",
                LeftArrowIcon{ }
            }
            h2{
                "Regler"
            }
            p {
                class: "text-xs",
                "Treman är ett dryckesspel som går ut på att slå tärningar för att bestämma vem som dricker och hur mycket."
            }
            h3 {
                "Vem är treman?"
            }
            p {
                class: "text-xs",
                "Börja med att utse en spelare till 'treman', vem som är treman kommer att ändras under spelets gång. Bara en person kan vara treman i taget."
            }
            p {
                class: "text-xs",
                "Följande leder till att en ny spelare blir 'treman':"
            }
            ul {
                class: "list-disc text-xs mx-4",
                li {"Om en ny person går med i spelet är hen nu treman."}
                li {"Om en person lämnar bordet och kommer tillbaka (t.ex. går på toa eller hämtar mer dricka) är hen nu treman."}
                li {"Skulle treman lämna bordet, blir föregående person treman igen."}
                li {"Vissa tärningsslag kan också resultera i en ny treman."}
            }
            p {
                class: "text-xs",
                "När en ny person blir treman skålar nya och gamla treman och tar vars en klunk."
            }
            h3 {
                "Sen då?"
            }
            p {
                class: "text-xs",
                "När ni valt treman kan spelet börja. En spelare slår tärningarna och ropar ut de regler som de slagit. En regel kan påverka en eller flera personer, se listan nedan, när alla påverkade har slutfört regeln slås tärningarna på nytt. Samma spelare fortsätter alltid att slå tärningarna tills dess att de slår 'Ingenting', då skickas tärningarna vidare medsols."
            }
            h3 {
                "Regellista"
            }
            for r in rules {
                DisplayRule { rule: r.clone() }
            }
        }
    }
}
