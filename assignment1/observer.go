package main

import "fmt"

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

type MangaSeries struct {
	Observers   []Observer
	Title       string
	Chapter     int
	IsAvailable bool
}

func NewMangaSeries(title string) *MangaSeries {
	return &MangaSeries{
		Title: title,
	}
}

// UpdateAvailability updates the availability of a new chapter.
func (m *MangaSeries) UpdateAvailability() {
	fmt.Printf("New chapter of %s is released (Chapter %d)\n", m.Title, m.Chapter)
	m.Chapter++
	m.IsAvailable = true
	m.NotifyAll()
}

func (m *MangaSeries) Register(observer Observer) {
	m.Observers = append(m.Observers, observer)
}

func (m *MangaSeries) Deregister(observer Observer) {
	for i, o := range m.Observers {
		if o == observer {
			m.Observers = append(m.Observers[:i], m.Observers[i+1:]...)
			break
		}
	}
}

func (m *MangaSeries) NotifyAll() {
	for _, observer := range m.Observers {
		observer.Update(m.Title, m.Chapter)
	}
}

type Observer interface {
	Update(title string, chapter int)
}

type Subscriber struct {
	ID string
}

func (s *Subscriber) Update(title string, chapter int) {
	fmt.Printf("Sending notification to subscriber %s: New chapter of %s (Chapter %d) is released\n", s.ID, title, chapter)
}

func main() {
	someManga := NewMangaSeries("Some Manga")

	subscriber1 := &Subscriber{ID: "subscriber1@gmail.com"}
	subscriber2 := &Subscriber{ID: "subscriber2@gmail.com"}

	someManga.Register(subscriber1)
	someManga.Register(subscriber2)

	someManga.UpdateAvailability()
}
