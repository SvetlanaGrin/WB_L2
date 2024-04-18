package repository

import (
	"dev11/internal/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type EventPostgres struct {
	db *sqlx.DB
}

type Event interface {
	CreateEvent(event entities.Event) (int, int, error)
	DeleteEvent(id int) (int, error)
	UpdateEvent(input entities.EventUpdate) (int, error)
	EventsForDay(day string) ([]entities.Event, int, error)
	EventsForWeek(day string, dayWeek string) ([]entities.Event, int, error)
	EventsForMonth(day string, dayMonth string) ([]entities.Event, int, error)
}

func NewEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{db: db}
}

func (e *EventPostgres) CreateEvent(event entities.Event) (int, int, error) {
	var id int = 1
	query := fmt.Sprintf("INSERT INTO events (event,day) values ($1, $2) RETURNING id")

	row := e.db.QueryRow(query, event.EventOne, event.Day)
	if err := row.Scan(&id); err != nil {
		return 0, 500, err
	}
	return id, 200, nil
}

func (e *EventPostgres) DeleteEvent(id int) (int, error) {
	querty := fmt.Sprintf("DELETE FROM Events WHERE id=$1")
	_, err := e.db.Exec(querty, id)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
func (e *EventPostgres) UpdateEvent(input entities.EventUpdate) (int, error) {
	id := input.Id

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.EventOne != "" {
		setValues = append(setValues, fmt.Sprintf("event=$%d", argId))
		args = append(args, input.EventOne)
		argId++
	}

	if input.Day != "" {
		setValues = append(setValues, fmt.Sprintf("day=$%d", argId))
		args = append(args, input.Day)
		argId++
	}
	var id1 int = 0
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE events tl SET %s WHERE tl.id=$%d RETURNING tl.id", setQuery, argId)
	args = append(args, id)
	logrus.Debugf("updateQuery:%s", query)
	logrus.Debugf("args:%s", args)
	row := e.db.QueryRow(query, args...)
	if err := row.Scan(&id1); err != nil {
		return 500, err
	}
	_, err := e.db.Exec(query, args...)

	if err != nil {
		return 500, err
	}
	return 200, nil
}
func (e *EventPostgres) EventsForDay(day string) ([]entities.Event, int, error) {
	var events []entities.Event
	query := fmt.Sprintf("SELECT * FROM events tl WHERE tl.day= $1")
	err := e.db.Select(&events, query, day)
	if err != nil {
		return events, 500, err
	}
	return events, 200, nil
}
func (e *EventPostgres) EventsForWeek(day string, dayWeek string) ([]entities.Event, int, error) {
	var events []entities.Event

	query := fmt.Sprintf("SELECT * FROM events tl WHERE tl.day>= $1 and tl.day<$2")
	err := e.db.Select(&events, query, day, dayWeek)
	if err != nil {
		return events, 500, err
	}
	return events, 200, nil
}
func (e *EventPostgres) EventsForMonth(day string, dayMonth string) ([]entities.Event, int, error) {
	var events []entities.Event
	query := fmt.Sprintf("SELECT * FROM events tl WHERE tl.day>= $1 and tl.day<=$2")
	err := e.db.Select(&events, query, day, dayMonth)
	if err != nil {
		return events, 500, err
	}
	return events, 200, nil
}
