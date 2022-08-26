package state

import (
	"changeme/utils"
	"log"
	"sync"
)

var (
	state = &AppState{
		Downloads: []DownloadInfo{},
	}
	stateMu     = &sync.Mutex{}
	subscribers = map[string]*subscriberInfo{}
	subMu       = &sync.Mutex{}
)

func GetState() *AppState {
	stateMu.Lock()
	defer stateMu.Unlock()
	return state
}

func Subscribe(cb subscriberCallback) (unsubscribe func()) {
	id := utils.GenUUID()
	sub := &subscriberInfo{
		id:       id,
		callback: cb,
	}

	subMu.Lock()
	defer subMu.Unlock()

	subscribers[id] = sub

	return func() {
		subMu.Lock()
		defer subMu.Unlock()

		delete(subscribers, id)
	}
}

func Dispatch(ac Action) error {
	stateMu.Lock()
	defer stateMu.Unlock()

	newState, err := reducer(state, ac)
	if err != nil {
		return err
	}

	state = newState
	notifySubscribers(newState)

	return nil
}

func notifySubscribers(state *AppState) {
	subMu.Lock()
	defer subMu.Unlock()

	for _, sub := range subscribers {
		if err := sub.callback(state); err != nil {
			log.Printf("Error: %v", err)
		}
	}
}
