package server

import (
	"L2/develop/dev11/model"
	"L2/develop/dev11/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	eventStorage service.EventStorage
}

func NewHandler() *Handler {
	return &Handler{
		eventStorage: service.NewEventStorage(),
	}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.EventsForDay)
	mux.HandleFunc("/events_for_week", h.EventsForWeek)
	mux.HandleFunc("/events_for_month", h.EventsForMonth)
}

// функция вывода ответа

func (h *Handler)ResultResponse(w http.ResponseWriter, m string ,res [] *model.Event) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.MarshalIndent(struct {
		Message string `json:"message"`
		Result []*model.Event `json:"result"`
	}{
		m,
		res,
	}, "", "\t")
	_,err:=w.Write(data)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

// проверка пользовательских данных

func (h *Handler) decodeJSON(r *http.Request) (*model.Event, error) {
	var event model.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}

	if event.UserID < 1 || event.EventID < 1 {
		return nil, errors.New("eventID or userID should pe positive")
	}

	return &event, nil
}



// функция вывода ошибки

func (h *Handler) errorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.MarshalIndent(struct {
		Error string `json:"error"`
	}{
		err.Error(),
	}, "", "\t")

	w.WriteHeader(status)
	w.Write(data)
}

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost{
		h.errorResponse(w,fmt.Errorf("incorrect method: %v;", "expected POST"),http.StatusBadRequest)
		return
	}

	 newEvent,err:=h.decodeJSON(r)
	if err != nil {
		h.errorResponse(w, fmt.Errorf("error while decoding input value: %v", err), http.StatusBadRequest)
		return
	}

	h.eventStorage.CreateEvent(newEvent)
	h.ResultResponse(w,"event has been created!",[]*model.Event{newEvent})

}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost{
		h.errorResponse(w,fmt.Errorf("incorrect method: %v;", "expected POST"),http.StatusBadRequest)
		return
	}

	newEvent,err:=h.decodeJSON(r)
	if err != nil {
		h.errorResponse(w, fmt.Errorf("error while decoding input value: %v", err), http.StatusBadRequest)
		return
	}

	h.eventStorage.UpdateEvent(newEvent)
	h.ResultResponse(w,"event has been updated!",[]*model.Event{newEvent})

}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost{
		h.errorResponse(w,fmt.Errorf("incorrect method: %v;", "expected POST"),http.StatusBadRequest)
		return
	}

	newEvent,err:=h.decodeJSON(r)
	if err != nil {
		h.errorResponse(w, fmt.Errorf("error while decoding input value: %v", err), http.StatusBadRequest)
		return
	}

	h.eventStorage.DeleteEvent(newEvent.EventID)
	h.ResultResponse(w,"event has been deleted!",[]*model.Event{newEvent})
}

func (h *Handler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet{
		h.errorResponse(w,fmt.Errorf("incorrect method: %v;", "expected GET"),http.StatusBadRequest)
		return
	}
	// парсим id и дату из запроса
	id,_:= strconv.Atoi(r.URL.Query().Get("user_id"))
	dateStr := r.URL.Query().Get("date")

	date, err := time.Parse(model.Layout, dateStr)
	if err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	events,err:=h.eventStorage.GetEventsForDay(id,date)
	if err != nil {
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.ResultResponse(w,"event for day: ",events)

}
func (h *Handler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet{
		h.errorResponse(w,fmt.Errorf("incorrect method: %v;", "expected GET"),http.StatusBadRequest)
		return
	}
	// парсим id и дату из запроса
	id,_:= strconv.Atoi(r.URL.Query().Get("user_id"))
	dateStr := r.URL.Query().Get("date")

	date, err := time.Parse(model.Layout, dateStr)
	if err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	events,err:=h.eventStorage.GetEventsForWeek(id,date)
	if err != nil {
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.ResultResponse(w,"",events)

}
func (h *Handler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet{
		h.errorResponse(w,fmt.Errorf("incorrect method: %v;", "expected GET"),http.StatusBadRequest)
		return
	}
	// парсим id и дату из запроса
	id,_:= strconv.Atoi(r.URL.Query().Get("user_id"))
	dateStr := r.URL.Query().Get("date")

	date, err := time.Parse(model.Layout, dateStr)
	if err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	events,err:=h.eventStorage.GetEventsForMonth(id,date)
	if err != nil {
		h.errorResponse(w, err, http.StatusServiceUnavailable)
		return
	}
	h.ResultResponse(w,"",events)

}
