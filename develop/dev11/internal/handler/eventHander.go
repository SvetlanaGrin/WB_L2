package handler

import (
	"dev11/internal/entities"
	"dev11/internal/handler/middleware"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, req *http.Request) {
	var input entities.Event
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		middleware.NewErrorResponse(w, 400, err)
		return
	}
	_, status, err := middleware.ValidTime(input.Day)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	id, status, err := h.services.CreateEvent(input)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	output, err := json.Marshal(map[string]interface{}{"result": id})
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
}
func (h *Handler) UpdateEvent(w http.ResponseWriter, req *http.Request) {
	var input entities.EventUpdate

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		middleware.NewErrorResponse(w, 400, err)
		return
	}
	status, err := h.services.UpdateEvent(input)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	output, err := json.Marshal(map[string]interface{}{"result": "Ok"})
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
}
func (h *Handler) DeleteEvent(w http.ResponseWriter, req *http.Request) {
	var input entities.EventId

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		middleware.NewErrorResponse(w, 400, err)
		return
	}
	status, err := h.services.DeleteEvent(input.Id)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	output, err := json.Marshal(map[string]interface{}{"result": "Ok"})
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
}
func (h *Handler) EventsForDay(w http.ResponseWriter, req *http.Request) {
	var events []entities.Event
	if err := req.ParseForm(); err != nil {
		middleware.NewErrorResponse(w, 400, err)
		return
	}

	day := req.Form.Get("day")

	_, status, err := middleware.ValidTime(day)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	events, status, err = h.services.EventsForDay(day)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	output, err := json.MarshalIndent(&events, "", "\t")
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
}
func (h *Handler) EventsForWeek(w http.ResponseWriter, req *http.Request) {
	var events []entities.Event
	if err := req.ParseForm(); err != nil {
		middleware.NewErrorResponse(w, 400, err)
		return
	}
	day := req.Form.Get("day")

	_, status, err := middleware.ValidTime(day)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	events, status, err = h.services.EventsForWeek(day)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}

	output, err := json.MarshalIndent(&events, "", "\t")
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
}
func (h *Handler) EventsForMonth(w http.ResponseWriter, req *http.Request) {
	var events []entities.Event
	if err := req.ParseForm(); err != nil {
		middleware.NewErrorResponse(w, 400, err)
		return
	}
	day := req.Form.Get("day")

	_, status, err := middleware.ValidTime(day)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}
	events, status, err = h.services.EventsForMonth(day)
	if err != nil {
		middleware.NewErrorResponse(w, status, err)
		return
	}

	output, err := json.MarshalIndent(&events, "", "\t")
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		middleware.NewErrorResponse(w, 500, err)
		return
	}
}
