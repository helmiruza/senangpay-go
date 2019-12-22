package senangpay

import (
  "fmt"
	"crypto/md5"
	"encoding/hex"
	"net/url"
  models "github.com/helmiruza/senangpay-go/models"
)

func generate_hash(secretKey string, detail string, amount string, orderId string) (string) {
	str := ""
	str += secretKey
	str += detail
	str += amount
	str += orderId

	hasher := md5.New()
  hasher.Write([]byte(str))
  return hex.EncodeToString(hasher.Sum(nil))
}

func CreatePaymentLink(data models.Payment) (models.Payment) {
	amount := fmt.Sprintf("%.2f", data.Amount)

	Url, err := url.Parse(URL)
	if err != nil {
    panic(err)
  }
  Url.Path += fmt.Sprintf("/payment/%s", MERCHANTID)

  parameters := url.Values{}
  parameters.Add("detail", data.Detail)
  parameters.Add("amount", amount)
  parameters.Add("order_id", data.OrderId)
	parameters.Add("name", data.Name)
	parameters.Add("email", data.Email)
	parameters.Add("phone", data.Phone)

	hashSign := generate_hash(SECRETKEY, data.Detail, amount, data.OrderId)
	parameters.Add("hash", hashSign)

  Url.RawQuery = parameters.Encode()

	data.Hash = hashSign
	data.Link = Url.String()

	return data
}
