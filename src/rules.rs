#[derive(Debug, Clone, Copy)]
pub enum RuleTrigger{
    Sum(u8),
    Pair(u8, u8),
    Single(u8)
}

impl RuleTrigger {
    pub fn inc(self, right: bool) -> Self {
        match self {
            Self::Sum(v) => Self::Sum(if v < 12 {
                v+1 
            }else {
                2
            }),
            Self::Pair(mut a, mut b) => {
                if !right {
                    a%=6;
                    a+=1;
                } else {
                    b%=6;
                    b+=1;
                }
                Self::Pair(a,b)
            }
            Self::Single(mut v) => Self::Single({
                v%=6;
                v+1
            }),
        }
    }
}

pub enum Rule {
    Base {
        trigger: RuleTrigger,
        name: &'static str,
        desc: &'static str,
    },
    User {
        trigger: RuleTrigger,
        name: String,
    }
}

impl Rule {
    pub const BASE: &[Self] = &[];
}
