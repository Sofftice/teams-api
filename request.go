package teamsapi

import (
	"context"
	"net/http"
)
func (c *Client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)

	return c.HTTPClient.Do(req)
}