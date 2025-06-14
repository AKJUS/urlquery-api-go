package urlquery

import (
	"fmt"
	"io"
)

func GetResource(report_id string, hash string) {
	DefaultClient.GetResource(report_id, hash)
}

func (api Client) GetResource(report_id string, hash string) ([]byte, error) {

	endpoint := fmt.Sprintf("/public/v1/report/%s/resource/%s", report_id, hash)
	resp, err := api.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = handleResponseError(resp)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, err
}
