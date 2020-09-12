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

func (r *walletRepository) UpdateMerchantWallet(c context.Context, merchantID uuid.UUID, amount float64) error {
	payload := map[string]interface{}{"amount": amount}
	log.Println(payload)
	reqBody, err := json.Marshal(payload)
	var URL string
	URL = fmt.Sprintf("%s/api/v1/transactions/users/%s/credit", r.APIURL, merchantID)

	req, err := http.NewRequest("PUT", URL, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(req.Body)
	log.Println(string(body))
	if err != nil {
		return err
	}
	// client := &http.Client{}
	// res, err := client.Do(req)
	// log.Println(URL, res.Status, res.StatusCode)
	// if err != nil {
	// 	return err
	// }
	// defer res.Body.Close()

	// _, err = ioutil.ReadAll(res.Body)
	return err
}
