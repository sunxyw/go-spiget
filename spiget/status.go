package spiget

import "context"

type StatusService service

type Status struct {
	Server struct {
		Name string `json:"name"`
		Mode string `json:"mode"`
	} `json:"server"`

	Fetch struct {
		Start       int64  `json:"start"`
		StartString string `json:"startString"`
		End         int64  `json:"end"`
		Active      bool   `json:"active"`
		Page        struct {
			Amount int `json:"amount"`
			Index  int `json:"index"`
			Item   struct {
				Index  int    `json:"index"`
				Status string `json:"status"`
			} `json:"item"`
		} `json:"page"`
	} `json:"fetch"`

	RestFetch struct {
		Start       int64  `json:"start"`
		StartString string `json:"startString"`
		End         int64  `json:"end"`
		Active      bool   `json:"active"`
		N           struct {
			Num   int `json:"num"`
			Start int `json:"start"`
			End   int `json:"end"`
			Max   int `json:"max"`
			Index int `json:"index"`
		} `json:"n"`
	} `json:"restFetch"`

	Existence struct {
		Start       int64  `json:"start"`
		StartString string `json:"startString"`
		End         int64  `json:"end"`
		Active      bool   `json:"active"`
		Document    struct {
			Amount   int `json:"amount"`
			Suspects int `json:"suspects"`
			Index    int `json:"index"`
			ID       int `json:"id"`
		} `json:"document"`
	} `json:"existence"`
}

type Stats struct {
	Resources        int `json:"resources"`
	Authors          int `json:"authors"`
	Categories       int `json:"categories"`
	ResourceUpdates  int `json:"resource_updates"`
	ResourceVersions int `json:"resource_versions"`
	Reviews          int `json:"reviews"`
}

type StatusResponse struct {
	Status *Status `json:"status"`
	Stats  *Stats  `json:"stats"`
}

func (sr *StatusResponse) String() string {
	return Stringify(sr)
}

// Get the API status
//
// Spiget API docs: https://spiget.org/documentation/#!/status/get_status
func (s *StatusService) Get(ctx context.Context) (*StatusResponse, *Response, error) {
	req, err := s.client.NewRequest("GET", "status", nil)
	if err != nil {
		return nil, nil, err
	}

	var statusResponse *StatusResponse
	resp, err := s.client.Do(ctx, req, &statusResponse)
	if err != nil {
		return nil, resp, err
	}

	return statusResponse, resp, nil
}
