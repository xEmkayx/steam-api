package model

type AchievementPercentagesResponse struct {
	GlobalAchievementPercentagesWrapper GlobalAchievementPercentages `json:"achievementpercentages" xml:"achievementpercentages"`
}

type GlobalAchievementPercentages struct {
	Achievements []GlobalAchievementPercentagesAchievement `json:"achievements" xml:"achievements>achievement"`
}

type GlobalAchievementPercentagesAchievement struct {
	Name    string  `json:"name" xml:"name"`
	Percent float64 `json:"percent" xml:"percent"`
}
