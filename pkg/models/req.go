package models

type AddMoneyReq struct {
	UserID int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}

type TransferMoneyReq struct {
	From   int64   `json:"from_user_id"`
	To     int64   `json:"to_user_id"`
	Amount float64 `json:"amount_to_transfer"`
}
