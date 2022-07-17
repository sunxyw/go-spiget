package spiget

// Icon represents a icon, usually user avatar.
type Icon struct {
	URL  string `json:"url,omitempty"`
	Data string `json:"data,omitempty"`
	Info string `json:"info,omitempty"`
	Hash string `json:"hash,omitempty"`
}

// File represents a file.
type File struct {
	Type        string  `json:"type,omitempty"`
	URL         string  `json:"url,omitempty"`
	Size        float64 `json:"size,omitempty"`
	SizeUnit    string  `json:"sizeUnit,omitempty"`
	ExternalUrl string  `json:"externalUrl,omitempty"`
}

// Rating represents a rating.
type Rating struct {
	Count   int     `json:"count,omitempty"`
	Average float64 `json:"average,omitempty"`
}

// Version represents a version.
type Version struct {
	ID   int    `json:"id,omitempty"`
	UUID string `json:"uuid,omitempty"`
}

// Update represents an update.
type Update struct {
	ID          int    `json:"id,omitempty"`
	Resource    int    `json:"resource,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Date        int    `json:"date,omitempty"`
	Likes       int    `json:"likes,omitempty"`
}

// Review represents a review.
type Review struct {
	ID       int    `json:"id,omitempty"`
	Resource int    `json:"resource,omitempty"`
	Author   Author `json:"author,omitempty"`
	Rating   Rating `json:"rating,omitempty"`
	Message  string `json:"message,omitempty"`
	Version  string `json:"version,omitempty"`
	Date     int    `json:"date,omitempty"`
}
