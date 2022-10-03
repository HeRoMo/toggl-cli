package toggl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HeRoMo/toggl-cli/toggl/model"
)

type toggl struct {
	userAgent string
	token     string
}

func Client(token string) *toggl {
	return &toggl{
		userAgent: "Toggl-CLI",
		token:     token,
	}
}

func (t *toggl) Reports(ctx context.Context, workspaceID int, since string, until string) (*model.DetailReport, error) {
	url := fmt.Sprintf(
		"https://api.track.toggl.com/reports/api/v2/details?workspace_id=%d&since=%s&until=%s&user_agent=%s",
		workspaceID,
		since,
		until,
		t.userAgent,
	)
	return getReport[model.DetailReport](ctx, url, t.token)
}

func (t *toggl) Summary(ctx context.Context, workspaceID int, since string, until string) (*model.SummaryReport, error) {
	url := fmt.Sprintf(
		"https://api.track.toggl.com/reports/api/v2/summary?workspace_id=%d&since=%s&until=%s&user_agent=%s",
		workspaceID,
		since,
		until,
		t.userAgent,
	)
	return getReport[model.SummaryReport](ctx, url, t.token)
}

func getReport[T model.ApiResponse](ctx context.Context, url string, token string) (*T, error) {
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.SetBasicAuth(token, "api_token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("ERROR")
		}
	}()

	byteArray, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s: %s", resp.Status, byteArray)
	}

	var report T
	if err := json.Unmarshal(byteArray, &report); err != nil {
		return nil, err
	}
	return &report, nil
}
