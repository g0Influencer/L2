package service

import (
	"L2/develop/dev11/model"
	"sync"
	"time"
)

//интерфейс календаря

type EventStorage interface {
	CreateEvent(event* model.Event)
	UpdateEvent(event * model.Event)
	DeleteEvent(eventID int)
	GetEventsForDay(userID int, date time.Time) ([]*model.Event,error)
	GetEventsForWeek(userID int, date time.Time) ([]*model.Event,error)
	GetEventsForMonth(userID int, date time.Time) ([]*model.Event,error)
}

// хранилище пользовательских запросов

type Storage struct {
	mu     *sync.Mutex
	idCount int
	events map[int]*model.Event
}
// инициализируем хранилище

func NewEventStorage () *Storage{
	return &Storage{
		events: make(map[int]*model.Event),
		idCount: 0,
		mu: new(sync.Mutex),
	}
}

// создание нового события

func (s *Storage) CreateEvent(event* model.Event){ // посмотреть,нужно ли проверять id на наличие в хранилище
	s.mu.Lock()
	event.EventID = s.idCount
	s.events[event.EventID] = event
	s.idCount++
	s.mu.Unlock()
}
// обновить событие

func (s *Storage) UpdateEvent(event * model.Event){
	s.mu.Lock()
	s.events[event.EventID] = event
	s.mu.Unlock()
}
// удалить событие

func (s *Storage) DeleteEvent(eventID int){
	s.mu.Lock()
	delete(s.events, eventID)
	s.mu.Unlock()
}

// вывести все события в этот день

func (s *Storage) GetEventsForDay(userID int, date time.Time) ([]*model.Event,error){
	eventsForDay:=make([]*model.Event,0)
	s.mu.Lock()

	for _,ev:=range s.events{
		if ev.Date.Year() == date.Year() && ev.Date.Month() == date.Month() && ev.Date.Day() == date.Day() && ev.UserID == userID{
			eventsForDay = append(eventsForDay,ev)
		}
	}
	s.mu.Unlock()
	return eventsForDay,nil
}
// вывести все события за неделю

func (s *Storage) GetEventsForWeek(userID int, date time.Time) ([]*model.Event,error){
	eventsForWeek:=make([]*model.Event,0)
	s.mu.Lock()

	for _,ev:=range s.events{
		y1,w1:=ev.Date.ISOWeek()
		y2,w2:=date.ISOWeek()
		if y1 == y2 && w1 == w2 && ev.UserID == userID{
			eventsForWeek = append(eventsForWeek,ev)
		}
	}
	s.mu.Unlock()

	return eventsForWeek,nil
}

// вывести все события за месяц

func (s *Storage) GetEventsForMonth(userID int, date time.Time) ([]*model.Event,error){
	eventsForMonth:=make([]*model.Event,0)
	s.mu.Lock()

	for _,ev:=range s.events{
		if ev.Date.Year() == date.Year() && ev.Date.Month() == date.Month()&& ev.UserID == userID{
			eventsForMonth = append(eventsForMonth,ev)
		}
	}
	s.mu.Unlock()
	return eventsForMonth,nil
}
