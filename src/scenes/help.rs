use crate::rules::*;
use dioxus::prelude::*;

#[inline_props]
fn DisplayTrigger<'a>(cx: Scope, trigger: &'a RuleTrigger) -> Element {
    match trigger {
        RuleTrigger::Sum(v) => render! {
            span {
                span {
                    class: "dice",
                    "0+0"
                }
                span { "={v}" }
            }
        },
        RuleTrigger::Pair(a, b) => render! {
            span {
                class: "dice",
                "{a}&{b}"
            }
        },
        RuleTrigger::Single(v) => render! {
            span {
                class: "dice",
                "{v}"
            }
        },
        RuleTrigger::Treman => render! {
            span {
                class: "dice",
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
                class: "p-4",
                p {
                    class: "text-xs",
                    DisplayTrigger { trigger: trigger }
                    span { " {name}" }
                }
                p {class: "text-xs", "{desc}"}
            }
        },
        Rule::User { trigger, name } => render! {
            p {
                class: "text-xs",
                DisplayTrigger { trigger: trigger }
                span { " {name}" }
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
        for r in rules {
            DisplayRule { rule: r.clone() }
        }
    }
}
