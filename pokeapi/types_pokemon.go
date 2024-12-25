package pokeapi

//RespPokemon
type RespPokemon struct {
    Name           string    `json:"name"`
    Height         int       `json:"height"`
    Weight         int       `json:"weight"`
    Stats          []Stat    `json:"stats"`
    Types          []Type    `json:"types"`
    BaseExperience int       `json:"base_experience"`
}

type Stat struct {
    BaseStat int      `json:"base_stat"`
    Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
    Name string `json:"name"`
}

type Type struct {
    Type TypeInfo `json:"type"`
}

type TypeInfo struct {
    Name string `json:"name"`
}
