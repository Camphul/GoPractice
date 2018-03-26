//Much credit to https://github.com/tmrts/go-patterns/tree/master/behavioral/observer
package main

import (
	"log"
	"time"
	"fmt"
)

type (

	Event struct {
		Payload []byte
	}

	Observer interface {
		OnNotify(Event)
	}

	Notifier interface {
		Register(Observer)
		Deregister(Observer)
		Notify(Event)
	}
)

type (
	//Observes Event interfaces
	eventObserver struct {
		id int
	}

	//Notifies observers
	eventNotifier struct {
		//Map of observers and the event listen for
		observers map[Observer]struct{}
	}
)

//OnNotify impl uses a unique copy of event for every observer
func (observer *eventObserver) OnNotify(e Event) {
	log.Printf("Observer %v received event: %v\n", observer.id, e)
}

func (notifier *eventNotifier) Register(observer Observer) {
	notifier.observers[observer] = struct{}{}
}

func (notifier *eventNotifier) Deregister(observer Observer) {
	delete(notifier.observers, observer)
}

func (notifier *eventNotifier) Notify(event Event) {
	for o := range notifier.observers {
		o.OnNotify(event)
	}
}

func craftPayload(t time.Time) []byte {
	return []byte(fmt.Sprintf("Current unix nano timestamp: %v/n",t.UnixNano()))
}

func main() {
	notifier := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// Register a couple of observers.
	notifier.Register(&eventObserver{id: 1})
	notifier.Register(&eventObserver{id: 2})

	// A simple loop publishing the current Unix timestamp to observers.
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <- stop:
			return
		case t := <-tick:
			notifier.Notify(Event{Payload: craftPayload(t)})
		}
	}
}

