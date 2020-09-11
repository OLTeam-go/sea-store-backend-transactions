package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/OLTeam-go/sea-store-backend-transactions/models"
	"github.com/google/uuid"
)

func validateResponseItems(ids []uuid.UUID, items []*models.Item) (bool, error) {
	m := make(map[string]bool)
	var quantityZero []string
	for _, item := range items {
		if item.Quantity > 0 {
			m[item.ID.String()] = true
		} else {
			quantityZero = append(quantityZero, item.ID.String())
		}
	}
	if len(quantityZero) != 0 {
		return false, fmt.Errorf("%s sold out", strings.Join(quantityZero, ", "))
	}
	var notFound []string
	for _, id := range ids {
		if !m[id.String()] {
			notFound = append(notFound, id.String())
		}
	}
	if len(notFound) != 0 {
		return false, fmt.Errorf("%s not found", strings.Join(notFound, ", "))
	}
	return true, nil
}

func (r *itemRepository) FetchByIDs(c context.Context, ids []uuid.UUID) ([]*models.Item, error) {
	payload := map[string]interface{}{"ids": ids}
	reqBody, err := json.Marshal(payload)
	URL := fmt.Sprintf("%s/api/items", r.APIURL)
	res, err := http.Post(URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	type responseBody struct {
		Status int            `json:"status"`
		Data   []*models.Item `json:"data"`
	}
	body, err := ioutil.ReadAll(res.Body)
	var result responseBody
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if ok, err := validateResponseItems(ids, result.Data); !ok {
		return nil, err
	}

	return result.Data, nil
}

func (r *itemRepository) UpdateQuantity(c context.Context, ids []uuid.UUID, action string) error {
	payload := map[string]interface{}{"ids": ids}
	reqBody, err := json.Marshal(payload)
	var URL string
	switch action {
	case "sold":
		URL = fmt.Sprintf("%s/api/items/sold", r.APIURL)
	case "available":
		URL = fmt.Sprintf("%s/api/items/available", r.APIURL)
	default:
		URL = fmt.Sprintf("%s/api/items/sold", r.APIURL)
	}
	res, err := http.Post(URL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	type responseBody struct {
		Status int            `json:"status"`
		Data   []*models.Item `json:"data"`
	}
	body, err := ioutil.ReadAll(res.Body)
	var result responseBody
	err = json.Unmarshal(body, &result)
	return err
}
