package senangpay

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
