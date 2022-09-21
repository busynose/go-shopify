package goshopify

import (
	"fmt"
	"net/http"
	"time"
)

const eventsBasePath = "events"
const eventsResourceName = "events"

// EventService is an interface for interfacing with the events endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/event
type EventService interface {
	List(interface{}) ([]Event, error)
	ListWithPagination(interface{}) ([]Event, *Pagination, error)
}

// EventServiceOp handles communication with the event related methods of the
// Shopify API.
type EventServiceOp struct {
	client *Client
}

// Represents the result from the events.json endpoint
type EventsResource struct {
	Events []Event `json:"events"`
}

// A struct for all available event list options.
// See: https://help.shopify.com/api/reference/event#index
type EventListOptions struct {
	CreatedAtMax time.Time `url:"created_at_max,omitempty"`
	CreatedAtMin time.Time `url:"created_at_min,omitempty"`
	Fields       string    `url:"fields,omitempty"`
	Limit        int       `url:"limit,omitempty"`
	SinceID      int64     `url:"since_id,omitempty"`
	verb         string    `url:"verb,omitempty"`
	filter       string    `url:"filter,omitempty"`
}

// Event represents a Shopify event
type Event struct {
	ID          int    `json:"id,omitempty"`
	SubjectID   int    `json:"subject_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	SubjectType string `json:"subject_type,omitempty"`
	Verb        string `json:"verb,omitempty"`
	Body        string `json:"body,omitempty"`
	Message     string `json:"message,omitempty"`
	Author      string `json:"author,omitempty"`
	Description string `json:"description,omitempty"`
	Path        string `json:"path,omitempty"`
}

// List events
func (s *EventServiceOp) List(options interface{}) ([]Event, error) {
	events, _, err := s.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventServiceOp) ListWithPagination(options interface{}) ([]Event, *Pagination, error) {
	path := fmt.Sprintf("%s.json", eventsBasePath)
	resource := new(EventsResource)
	headers := http.Header{}

	headers, err := s.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}

	// Extract pagination info from header
	linkHeader := headers.Get("Link")

	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}

	return resource.Events, pagination, nil
}
