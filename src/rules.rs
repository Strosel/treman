#[derive(Debug, Clone, Copy, PartialEq)]
pub enum RuleTrigger {
    Sum(u8),
    Pair(u8, u8),
    Single(u8),
    Treman,
}

impl RuleTrigger {
    pub fn inc(self, right: bool) -> Self {
        match self {
            Self::Sum(v) => Self::Sum(if v < 12 { v + 1 } else { 2 }),
            Self::Pair(mut a, mut b) => {
                if !right {
                    a %= 6;
                    a += 1;
                } else {
                    b %= 6;
                    b += 1;
                }
                Self::Pair(a, b)
            }
            Self::Single(mut v) => Self::Single({
                v %= 6;
                v + 1
            }),
            Self::Treman => Self::Treman,
        }
    }

    pub fn check(&self, dice: &(u8, u8)) -> bool {
        match self {
            Self::Sum(v) => (dice.0 + dice.1) == *v,
            Self::Pair(a, b) => *dice == (*a, *b) || *dice == (*b, *a),
            Self::Single(v) => dice.0 == *v || dice.1 == *v,
            Self::Treman => (dice.0 == 3 || dice.1 == 3) && dice.0 != dice.1,
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub enum Rule {
    Base {
        trigger: RuleTrigger,
        name: &'static str,
        desc: &'static str,
    },
    User {
        trigger: RuleTrigger,
        name: String,
    },
}

impl Rule {
    pub fn check(&self, dice: &(u8, u8)) -> bool {
        match self {
            Self::Base { trigger, .. } | Self::User { trigger, .. } => trigger.check(dice),
        }
    }

    pub fn name(&self) -> &str {
        match self {
            Self::Base { name, .. } => name,
            Self::User { name, .. } => name,
        }
    }

    pub const BASE: &[Self] = &[
        Rule::Base {
            name: "Treman",
            desc: "Treman dricker",
            trigger: RuleTrigger::Treman,
        },
        Rule::Base {
            name: "Krig",
            desc: "Välj en annan spelare. Ni är nu i krig, dricker den ena så dricker bägge",
            trigger: RuleTrigger::Pair(1, 1),
        },
        Rule::Base {
            name: "Utmaning",
            desc:
                "Välj en annan spelare och slå vars en tärning. Den som slår högst blir ny treman.",
            trigger: RuleTrigger::Pair(1, 2),
        },
        Rule::Base {
            name: "En ferrari",
            desc: "Sist att låtsas köra bil dricker. (\"Dark humour\" variant finns)",
            trigger: RuleTrigger::Pair(1, 4),
        },
        Rule::Base {
            name: "Ny Treman",
            desc: "Grattis! Du är nu treman",
            trigger: RuleTrigger::Pair(3, 3),
        },
        Rule::Base {
            name: "Jag har aldrig sett...",
            desc: concat!(
                "Häfv resten av din enhet och skapa en ny regel.\n",
                "Eller dela ut 6+6 klunkar, (6 klunkar till två personer eller 12 till en person)",
            ),
            trigger: RuleTrigger::Pair(6, 6),
        },
        Rule::Base {
            name: "Dela ut 2+2 klunkar",
            desc: "Dela ut 2 klunkar till två personer eller 4 till en person",
            trigger: RuleTrigger::Pair(2, 2),
        },
        Rule::Base {
            name: "Dela ut 4+4 klunkar",
            desc: "Dela ut 4 klunkar till två personer eller 8 till en person",
            trigger: RuleTrigger::Pair(4, 4),
        },
        Rule::Base {
            name: "Dela ut 5+5 klunkar",
            desc: "Dela ut 5 klunkar till två personer eller 10 till en person",
            trigger: RuleTrigger::Pair(5, 5),
        },
        Rule::Base {
            name: "Seven ahead",
            desc: "Personen vänster om den som slår tärningarna dricker.",
            trigger: RuleTrigger::Sum(7),
        },
        Rule::Base {
            name: "Nine behind",
            desc: "Personen höger om den som slår tärningarna dricker.",
            trigger: RuleTrigger::Sum(9),
        },
        Rule::Base {
            name: "Finger på näsan",
            desc: "Sist att sätta fingret på näsan dricker.",
            trigger: RuleTrigger::Sum(11),
        },
    ];
}
