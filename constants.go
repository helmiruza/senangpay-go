package senangpay

import (
  "encoding/json"
  models "github.com/helmiruza/senangpay-go/models"
)

var ENVIRONMENT = ""
var MERCHANTID = ""
var SECRETKEY = ""
var URL = ""

const (
  ProductionUrl = "https://api.senangpay.my"
	StagingUrl = "https://sandbox.senangpay.my"
)

func Init(e string, f string, g string) {
	ENVIRONMENT = e
  MERCHANTID = f
  SECRETKEY = g

  if ENVIRONMENT == "production" {
    URL = ProductionUrl
  }

  if ENVIRONMENT == "staging" {
    URL = StagingUrl
  }
}

func objectFactory(objectType string) (interface{}) {
  switch (objectType) {
    case "collection":
      return models.CollectionResponse{}
    case "bill":
      return models.BillResponse{}
  }
  return nil
}

func response(body []byte, err error, objectType string) (interface{}, models.Error) {
  errorMessage := models.Error{}
  response := objectFactory(objectType)

  if err != nil {
    errorMessage.Error = models.ErrorDetail { Type: "unknown", Message: "unknown" }
    return response, errorMessage
  }

  json.Unmarshal(body, &errorMessage)

  if len(errorMessage.Error.Type) > 0 {
    return response, errorMessage
  }

  json.Unmarshal(body, &response)

  return response, errorMessage
}
