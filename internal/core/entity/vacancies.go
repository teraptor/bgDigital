package entity

import "time"

/**
все вакансии которые находятся в одном из статусов работы
*/

type ChannelIDEnum int

const (
	TelegramChannelIDEnum = 1
)

// вакансия
type Vacancies struct {
	ID        int           `bson:"id"`
	MessageID int           `bson:"message_id"`
	UserID    int           `bson:"user_id"`
	PartnerID int           `bson:"partner_id"`
	Message   string        `bson:"message"`
	CreatedAt time.Time     `bson:"created_at"`
	DeletedAt time.Time     `bson:"deleted_at"`
	ChannelID ChannelIDEnum `bson:"channel_id"`
}
