package models

type Problem struct {
	ID                 uint     `gorm:"primaryKey" json:"id"`
	Title              string   `json:"title"`
	TitleSlug          string   `json:"titleSlug"`
	FrontendQuestionId string   `json:"frontendQuestionId"`
	Difficulty         string   `json:"difficulty"`
	TopicTags          []string `gorm:"-" json:"topicTags"` // `gorm:"-"` skips DB column mapping
}
