use super::Scene;
use crate::components::*;
use crate::rules::*;
use dioxus::prelude::*;

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
            div {
                class: "flex flex-row h-[4vh] w-100 justify-between items-center",

                Link {
                    to: Scene::Game,
                    class: "icon",
                    LeftArrowIcon { }
                }

                a {
                    class: "icon",
                    href: "https://github.com/strosel/treman",
                    GitHubIcon { }
                }
            }

            h2{
                "Regler"
            }
            p {
                class: "text-xs",
                "Treman är ett dryckesspel som går ut på att slå tärningar för att bestämma vem som dricker och hur mycket."
            }
            h3 {
                "Vem är Treman?"
            }
            p {
                class: "text-xs",
                "Börja med att utse en spelare till 'Treman', vem som är Treman kommer att ändras under spelets gång. Bara en person kan vara Treman i taget."
            }
            p {
                class: "text-xs",
                "Följande leder till att en ny spelare blir 'Treman':"
            }
            ul {
                class: "list-disc text-xs mx-4",
                li {"Om en ny person går med i spelet är hen nu Treman."}
                li {"Om en person lämnar bordet och kommer tillbaka (t.ex. går på toa eller hämtar mer dricka) är hen nu Treman."}
                li {"Skulle Treman lämna bordet, blir föregående person Treman igen."}
                li {"Vissa tärningsslag kan också resultera i en ny Treman."}
            }
            p {
                class: "text-xs",
                "När en ny person blir Treman skålar nya och gamla Treman och tar vars en klunk."
            }
            h3 {
                "Sen då?"
            }
            p {
                class: "text-xs",
                "När ni valt Treman kan spelet börja. En spelare slår tärningarna och ropar ut de regler som de slagit. En regel kan påverka en eller flera personer, se listan nedan, när alla påverkade har slutfört regeln slås tärningarna på nytt. Samma spelare fortsätter alltid att slå tärningarna tills dess att de slår 'Ingenting', då skickas tärningarna vidare medsols."
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
