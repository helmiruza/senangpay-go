package senangpay

import (
  "fmt"
	"net/http"
  "io/ioutil"
  "bytes"
  "encoding/json"
  models "github.com/helmiruza/senangpay-go/models"
)

func CreatePaymentLink(data models.Payment) (string) {
  URL += fmt.Sprintf("/payment/%s", MERCHANTID)

  return URL
}
