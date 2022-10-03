package model

type ReportBase struct {
	TotalGrand      int               `json:"total_grand"`
	TotalBillable   string            `json:"total_billable"`
	TotalCount      int               `json:"total_count"`
	PerPage         int               `json:"per_page"`
	TotalCurrencies []TotalCurrencies `json:"total_currencies"`
}

type TotalCurrencies struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type ApiResponse interface {
	DetailReport | SummaryReport
}

type DetailReport struct {
	ReportBase
	Data []DetailReportItem `json:"data"`
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
