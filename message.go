package sugar

import (
	"errors"
	"time"
)

type Message struct {
	// Contents
}

type Subscription struct {
	ch chan Message

	Inbox chan Message
}

func (s *Subscription) Publish(msg Message) error {
	if _, ok := <-s.ch; !ok {
		return errors.New("Topic has been closed")
	}

	s.ch <- msg

	return nil
}

type Topic struct {
	Subscribers    []Session
	MessageHistory []Message
	ch             <-chan Message
}

func (t *Topic) Subscribe(uid uint64, name string) (Subscription, error) {
	// Get session and create one if it's the first
	s := Session{
		User: User{
			ID:   uid,
			Name: name,
		},
		Timestamp: time.Now(),
	}

	// Add session to the Topic & MessageHistory
	t.Subscribers = append(t.Subscribers, s)
	// Create a subscription
	return Subscription{}, nil
}

func (t *Topic) Unsubscribe(Subscription) error {
	// Implementation
	return nil
}

func (t *Topic) Delete() error {
	// Implementation
	return nil
}

type User struct {
	ID   uint64
	Name string
}

type Session struct {
	User      User
	Timestamp time.Time
}

var (
	gSubscription Subscription
	gTopic        Topic
	c             = make(chan Message)
)

func work() {

	gSubscription.ch = c
	gSubscription.Publish(Message{})

}

func hub() {
	for {
		select {
		case data := <-c:
			gTopic.MessageHistory = append(gTopic.MessageHistory, data)

		}
	}

}
