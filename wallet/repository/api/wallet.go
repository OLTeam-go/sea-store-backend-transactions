package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (r *walletRepository) UpdateMerchantWallet(c context.Context, merchantID uuid.UUID, amout float64) error {
	payload := map[string]interface{}{"amount": amout}
	reqBody, err := json.Marshal(payload)
	var URL string
	URL = fmt.Sprintf("%s/api/v1/transactions/users/%s/credit", r.APIURL, merchantID)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	res, err := client.Do(req)
	log.Println(URL, res.Status, res.StatusCode)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	return err
}
