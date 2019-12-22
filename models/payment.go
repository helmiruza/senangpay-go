package models

type Payment struct {
  Detail string `json:"detail"`
  Amount float `json:"amount"`
  OrderId string `json:"order_id"`
  Hash string `json:"hash"`
  Name string `json:"name"`
  Email string `json:"email"`
  Phone string `json:"phone"`
}

type PaymentResponse struct {
  StatusId string `json:"statud_id"`
  OrderId string `json:"order_id"`
  Msg string `json:"msg"`
  TransactionId string `json:"transaction_id"`
  Hash string `json:"hash"`
}
