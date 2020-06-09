package rest

type NewSession struct {
	SessionID string `json:"session_id"`
	IPAddress string `json:"ip_address"`
	Date time.Time `json:"date"`
}
