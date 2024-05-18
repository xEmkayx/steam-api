package model

type SchemaForGameResponse struct {
	SchemaForGameWrapper GameSchemaGame `json:"game" xml:"game"`
}

// Root structure
type GameSchemaGame struct {
	GameName           string                       `json:"gameName" xml:"gameName"`
	GameVersion        string                       `json:"gameVersion" xml:"gameVersion"`
	AvailableGameStats GameSchemaAvailableGameStats `json:"availableGameStats" xml:"availableGameStats"`
}

// Sub-structure for available game stats
type GameSchemaAvailableGameStats struct {
	Stats []GameSchemaStat `json:"stats" xml:"stats>stat"`
	// TODO: testen - welches der beiden?
	// Achievements []Achievement    `json:"achievements" xml:"achievements>achievement"`
	Achievements []GameSchemaAchievement `json:"achievements" xml:"achievements>achievement"`
}

// Stat structure
type GameSchemaStat struct {
	Name         string `json:"name" xml:"name"`
	DefaultValue int    `json:"defaultvalue" xml:"defaultvalue"`
	DisplayName  string `json:"displayName" xml:"displayName"`
}

// Achievement structure
type GameSchemaAchievement struct {
	Name         string `json:"name" xml:"name"`
	DefaultValue int    `json:"defaultvalue" xml:"defaultvalue"`
	DisplayName  string `json:"displayName" xml:"displayName"`
	Hidden       int    `json:"hidden" xml:"hidden"`
	Icon         string `json:"icon" xml:"icon"`
	IconGray     string `json:"icongray" xml:"icongray"`
}
