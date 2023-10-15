package steam

// Result defines the common result of web API returned.
type Result struct {
	Response interface{} `json:"response"`
}
