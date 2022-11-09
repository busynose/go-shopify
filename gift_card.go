package goshopify

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

const giftCardsBasePath = "gift_cards"

type GiftCardService interface {
	List(interface{}) ([]GiftCard, error)
	//Count(interface{}) (int, error)
	//Get(int64, interface{}) (*GiftCard, error)
	Create(GiftCard) (*GiftCard, error)
	//Update(GiftCard) (*GiftCard, error)
	//Delete(int64) error
}

type GiftCardServiceOp struct {
	client *Client
}

type GiftCard struct {
	ID             int64            `json:"id,omitempty"`
	Balance        *decimal.Decimal `json:"balance,omitempty"`
	Code           string           `json:"code,omitempty"`
	CreatedAt      *time.Time       `json:"created_at,omitempty"`
	UpdatedAt      *time.Time       `json:"updated_at,omitempty"`
	DisabledAt     *time.Time       `json:"disabled_at,omitempty"`
	InitialValue   *decimal.Decimal `json:"initial_value,omitempty"`
	ExpiresOn      string           `json:"expires_on,omitempty"`
	Currency       string           `json:"currency,omitempty"`
	LineItemID     int64            `json:"line_item_id,omitempty"`
	APIClientID    int64            `json:"api_client_id,omitempty"`
	UserID         int64            `json:"user_id,omitempty"`
	CustomerID     int64            `json:"customer_id,omitempty"`
	Note           string           `json:"note,omitempty"`
	TemplateSuffix string           `json:"template_suffix,omitempty"`
	LastCharacters string           `json:"last_characters,omitempty"`
	OrderID        int64           `json:"order_id,omitempty"`
}

type GiftCardResource struct {
	GiftCard *GiftCard `json:"gift_card"`
}

type GiftCardsResource struct {
	GiftCards []GiftCard `json:"gift_cards"`
}

// List giftCards
func (s *GiftCardServiceOp) List(options interface{}) ([]GiftCard, error) {
	giftCards, _, err := s.ListWithPagination(options)
	if err != nil {
		return nil, err
	}
	return giftCards, nil
}

func (s *GiftCardServiceOp) ListWithPagination(options interface{}) ([]GiftCard, *Pagination, error) {
	path := fmt.Sprintf("%s.json", giftCardsBasePath)
	resource := new(GiftCardsResource)
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

	return resource.GiftCards, pagination, nil
}

// Create a gift card
func (s *GiftCardServiceOp) Create(giftCard GiftCard) (*GiftCard, error) {
	path := fmt.Sprintf("%s.json", giftCardsBasePath)
	wrappedData := GiftCardResource{GiftCard: &giftCard}
	resource := new(GiftCardResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.GiftCard, err
}
