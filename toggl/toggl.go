package toggl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func (t *toggl) Reports(workspaceID int, since string, until string) (*DetailReport, error) {
	url := fmt.Sprintf(
		"https://api.track.toggl.com/reports/api/v2/details?workspace_id=%d&since=%s&until=%s&user_agent=%s",
		workspaceID,
		since,
		until,
		t.userAgent,
	)
	return getReport[DetailReport](url, t.token)
}

func (t *toggl) Summary(workspaceID int, since string, until string) (*SummaryReport, error) {
	url := fmt.Sprintf(
		"https://api.track.toggl.com/reports/api/v2/summary?workspace_id=%d&since=%s&until=%s&user_agent=%s",
		workspaceID,
		since,
		until,
		t.userAgent,
	)
	return getReport[SummaryReport](url, t.token)
}

type ApiResponse interface {
	DetailReport | SummaryReport
}

func getReport[T ApiResponse](url string, token string) (*T, error) {
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
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

	byteArray, _ := ioutil.ReadAll(resp.Body)

	var report T
	if err := json.Unmarshal(byteArray, &report); err != nil {
		log.Fatal(err)
	}
	return &report, nil
}

type DetailReportItem struct {
	ID              int      `json:"id"`
	PID             int      `json:"pid"`
	TID             int      `json:"tid"`
	UID             int      `json:"uid"`
	Description     string   `json:"description"`
	Start           string   `json:"start"`
	End             string   `json:"end"`
	Updated         string   `json:"updated"`
	Dur             int      `json:"dur"`
	User            string   `json:"user"`
	UseStop         bool     `json:"use_stop"`
	Client          string   `json:"client"`
	Project         string   `json:"project"`
	ProjectColor    string   `json:"project_color"`
	ProjectHexColor string   `json:"project_hex_color"`
	Task            string   `json:"task"`
	Billable        string   `json:"billable"`
	IsBillable      bool     `json:"is_billable"`
	Cur             string   `json:"cur"`
	Tags            []string `json:"tags"`
}

type TotalCurrencies struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type ReportBase struct {
	TotalGrand      int               `json:"total_grand"`
	TotalBillable   string            `json:"total_billable"`
	TotalCount      int               `json:"total_count"`
	PerPage         int               `json:"per_page"`
	TotalCurrencies []TotalCurrencies `json:"total_currencies"`
}

type DetailReport struct {
	ReportBase
	Data []DetailReportItem `json:"data"`
}

type SummaryReport struct {
	ReportBase
	Data []SummaryReportData `json:"data"`
}

type SummaryReportData struct {
	ID    int `json:"id"`
	Title struct {
		Project string `json:"project"`
		Client  string `json:"client"`
		User    string `json:"user"`
	}
	Time            int               `json:"time"`
	TotalCurrencies []TotalCurrencies `json:"total_currencies"`
	Items           []SummaryReportItem
}

type SummaryReportItem struct {
	Title struct {
		Project   string `json:"project"`
		Client    string `json:"client"`
		User      string `json:"user"`
		Task      string `json:"task"`
		TimeEntry string `json:"time_entry"`
	} `json:"title"`
	Time int    `json:"time"`
	Cur  string `json:"cur"`
	Sum  int    `json:"sum"`
	Rate int    `json:"rate"`
}
