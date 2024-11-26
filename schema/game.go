package schema

type Game struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"jogadores"`
	Kills      map[string]int `json:"kills"`
	Deaths     map[string]int `json:"mortes_por_meio"`
}

func LoadMeansOfDeath() map[string]string {
	return map[string]string{
		"MOD_UNKNOWN":        "Desconhecido",
		"MOD_SHOTGUN":        "Espingarda",
		"MOD_GAUNTLET":       "Manopla",
		"MOD_MACHINEGUN":     "Metralhadora",
		"MOD_GRENADE":        "Granada",
		"MOD_GRENADE_SPLASH": "Explosão de Granada",
		"MOD_ROCKET":         "Foguete",
		"MOD_ROCKET_SPLASH":  "Explosão de Foguete",
		"MOD_PLASMA":         "Plasma",
		"MOD_PLASMA_SPLASH":  "Explosão de Plasma",
		"MOD_RAILGUN":        "Railgun",
		"MOD_LIGHTNING":      "Raio",
		"MOD_BFG":            "BFG",
		"MOD_BFG_SPLASH":     "Explosão de BFG",
		"MOD_WATER":          "Afogamento",
		"MOD_SLIME":          "Lodo",
		"MOD_LAVA":           "Lava",
		"MOD_CRUSH":          "Esmagamento",
		"MOD_TELEFRAG":       "Telefrag",
		"MOD_FALLING":        "Queda",
		"MOD_SUICIDE":        "Suicídio",
		"MOD_TARGET_LASER":   "Laser",
		"MOD_TRIGGER_HURT":   "Dano por Trigger",
		"MOD_NAIL":           "Prego",
		"MOD_CHAINGUN":       "Metralhadora Giratória",
		"MOD_PROXIMITY_MINE": "Mina de Proximidade",
		"MOD_KAMIKAZE":       "Kamikaze",
		"MOD_JUICED":         "Juiced",
		"MOD_GRAPPLE":        "Grapple",
	}

}
