package service

import (
	"dev11/internal/entities"
	"dev11/internal/repository"
	"time"
)

type EventService struct {
	repo repository.Event
}

type Event interface {
	CreateEvent(event entities.Event) (int, int, error)
	DeleteEvent(id int) (int, error)
	UpdateEvent(input entities.EventUpdate) (int, error)
	EventsForDay(day string) ([]entities.Event, int, error)
	EventsForWeek(day string) ([]entities.Event, int, error)
	EventsForMonth(day string) ([]entities.Event, int, error)
}

func NewOrderService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (e *EventService) CreateEvent(event entities.Event) (int, int, error) {
	return e.repo.CreateEvent(event)
}
func (e *EventService) DeleteEvent(id int) (int, error) {
	return e.repo.DeleteEvent(id)
}
func (e *EventService) UpdateEvent(input entities.EventUpdate) (int, error) {
	return e.repo.UpdateEvent(input)
}
func (e *EventService) EventsForDay(day string) ([]entities.Event, int, error) {
	events, status, err := e.repo.EventsForDay(day)
	if err != nil {
		return events, status, err
	}
	layout := "2006-01-02T15:04:05Z"
	layoutDay := "2006-01-02"
	for i, elem := range events {
		t, err := time.Parse(layout, elem.Day)
		if err != nil {
			return events, 503, err
		}

		events[i].Day = t.Format(layoutDay)
	}
	return events, status, err
}
func (e *EventService) EventsForWeek(day string) ([]entities.Event, int, error) {
	week := time.Hour * 24 * 7
	layoutDay := "2006-01-02"
	t, _ := time.Parse(layoutDay, day)
	t = t.Add(week)
	events, status, err := e.repo.EventsForWeek(day, t.Format(layoutDay))
	if err != nil {
		return events, status, err
	}
	layout := "2006-01-02T15:04:05Z"
	for i, elem := range events {
		t, err := time.Parse(layout, elem.Day)
		if err != nil {
			return events, 503, err
		}

		events[i].Day = t.Format(layoutDay)
	}
	return events, status, err
}
func (e *EventService) EventsForMonth(day string) ([]entities.Event, int, error) {
	layoutDay := "2006-01-02"
	d, _ := time.Parse(layoutDay, day)
	t1 := time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(d.Year(), d.Month(), 29, 0, 0, 0, 0, time.UTC)
	events, status, err := e.repo.EventsForMonth(t1.Format(layoutDay), t2.Format(layoutDay))
	if err != nil {
		return events, status, err
	}
	layout := "2006-01-02T15:04:05Z"
	for i, elem := range events {
		t, err := time.Parse(layout, elem.Day)
		if err != nil {
			return events, 503, err
		}

		events[i].Day = t.Format(layoutDay)
	}
	return events, status, err
}
