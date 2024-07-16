package profiling // 导入 profiling 包

type Request struct {
	TransactionID string `json:"transaction_id"` // 交易 ID, JSON 标签 "transaction_id"
	PayLoad       []int  `json:"payload"`        // 负载, JSON 标签 "payload"
}

type Response struct {
	TransactionID string `json:"transaction_id"` // 交易 ID, JSON 标签 "transaction_id"
	Expression    string `json:"exp"`            // 表达式, JSON 标签 "exp"
}
