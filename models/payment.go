package models

type Payment struct {
  Detail string `json:"detail"`
  Amount float64 `json:"amount"`
  OrderId string `json:"order_id"`
  Hash string `json:"hash"`
  Name string `json:"name"`
  Email string `json:"email"`
  Phone string `json:"phone"`
  Link string `json:"link"`
}

type PaymentResponse struct {
  StatusId string `json:"statud_id"`
  OrderId string `json:"order_id"`
  Msg string `json:"msg"`
  TransactionId string `json:"transaction_id"`
  Hash string `json:"hash"`
}
