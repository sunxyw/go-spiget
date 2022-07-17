package spiget

// Icon represents a icon, usually user avatar.
type Icon struct {
	URL  string `json:"url,omitempty"`
	Data string `json:"data,omitempty"`
	Info string `json:"info,omitempty"`
	Hash string `json:"hash,omitempty"`
}
