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
	//Create(GiftCard) (*GiftCard, error)
	//Update(GiftCard) (*GiftCard, error)
	//Delete(int64) error
}

type GiftCardServiceOp struct {
	client *Client
}

type GiftCard struct {
	ID             int64            `json:"id"`
	Balance        *decimal.Decimal `json:"balance"`
	InitialValue   *decimal.Decimal `json:"initial_value"`
	CreatedAt      *time.Time       `json:"created_at"`
	UpdatedAt      *time.Time       `json:"updated_at"`
	DisabledAt     *time.Time       `json:"disabled_at"`
	ExpiresOn      *time.Time       `json:"expires_on"`
	Currency       string           `json:"currency"`
	LineItemID     int64            `json:"line_item_id"`
	APIClientID    int64            `json:"api_client_id"`
	UserID         int64            `json:"user_id"`
	CustomerID     int64            `json:"customer_id"`
	Note           string           `json:"note"`
	TemplateSuffix string           `json:"template_suffix"`
	LastCharacters string           `json:"last_characters"`
	OrderID        string           `json:"order_id"`
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
