package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mikelawson03/chores/internal/auth"
	"github.com/mikelawson03/chores/internal/domain"
	"github.com/mikelawson03/chores/internal/events"
)

func (cfg *apiCfg) handlerGetEvents(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthenticatedUser(r.Context())
	if err != nil {
		RespondWithError(w, domain.ErrUnauthorized)
		return
	}

	subscriberID, eventCh := cfg.App.Bus.Subscribe(
		user,
		events.SubscriberTypeClient,
	)

	defer cfg.App.Bus.Unsubscribe(subscriberID)

	fmt.Println("SSE client connected")

	w.Header().Set("Content-Type", "text/event-stream")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	flusher.Flush()

	for {
		select {
		case event, ok := <-eventCh:
			if !ok {
				return
			}

			data, err := json.Marshal(event)
			if err != nil {
				RespondWithError(w, domain.ErrInvalidRequest)
			}
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println("SSE client disconnected")
			return

		}
	}

}
