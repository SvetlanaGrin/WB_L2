package entities

type Event struct {
	Id       int    `db:"id"`
	EventOne string `json:"event" binding:"required" db:"event"`
	Day      string `json:"day" binding:"required" db:"day"`
}
type EventId struct {
	Id int `json:"id" binding:"required" `
}

type EventUpdate struct {
	Id       int    `json:"id" binding:"required" `
	EventOne string `json:"event"  `
	Day      string `json:"day" `
}
