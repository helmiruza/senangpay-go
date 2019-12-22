package senangpay

import (
  "fmt"
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  models "github.com/helmiruza/senangpay-go/models"
)

func CreatePayment(data models.Payment) (models.PaymentResponse, models.Error) {
  URL += fmt.Sprintf("/payment/%s", MERCHANTID)

  requestBody, _ := json.Marshal(data)

  client := &http.Client{}

  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
  req.SetBasicAuth(APIKEY, "")
  req.Header.Set("Content-type", "application/json")

  resp, _ := client.Do(req)
  body, err := ioutil.ReadAll(resp.Body)

  obj, err1 := response(body, err, "bill")
  objString, _ := json.Marshal(obj)
  k := models.BillResponse{}
  json.Unmarshal(objString, &k)

  return k, err1
}
