package qrtickets

// Venue - A physical location that hosts events
type Venue struct {
	Name    string `json:"name" datastore:",noindex"`
	Address string `json:"address" datatore:",noindex"`
	URL     string `json:"url" datastore:",noindex"`
}
