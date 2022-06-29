package observer

import (
	"fmt"

	"github.com/google/uuid"
)

func RunObserver() {
	p := webhooksPublisher{subscribers: make(map[uuid.UUID]Subscriber)}
	p.addSubscriber(ordersSubscriber{id: uuid.New()})

	p.notifySubscribers(Message{
		Type:    "orders/created",
		Payload: "payload order created",
	})
}

type Publisher interface {
	addSubscriber(subscriber Subscriber)
	removeSubscriber(subscriber Subscriber)
	notifySubscribers(message Message)
}

type Message struct {
	Type    string
	Payload string
}

type Subscriber interface {
	getID() uuid.UUID
	update(message Message)
}

type webhooksPublisher struct {
	subscribers map[uuid.UUID]Subscriber
}

func (p webhooksPublisher) addSubscriber(subscriber Subscriber) {
	p.subscribers[uuid.New()] = subscriber
}

func (p webhooksPublisher) removeSubscriber(subscriber Subscriber) {
	delete(p.subscribers, subscriber.getID())
}

func (p webhooksPublisher) notifySubscribers(message Message) {
	for _, subscriber := range p.subscribers {
		subscriber.update(message)
	}
}

type ordersSubscriber struct {
	id uuid.UUID
}

func (s ordersSubscriber) getID() uuid.UUID {
	return s.id
}

func (s ordersSubscriber) update(message Message) {
	if message.Type == "orders/created" {
		fmt.Println("updated with orders/created type")
		handleOrderCreated(message.Payload)
	}
}

func handleOrderCreated(payload string) {
	fmt.Println("handled order created")
}
