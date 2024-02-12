package main

import (
	"encoding/json"
)

type AutoGenerated2 struct {
	Query struct {
		Ids           []int    `json:"ids"`
		Preset        string   `json:"preset"`
		Dimensions    []string `json:"dimensions"`
		Metrics       []string `json:"metrics"`
		Sort          []string `json:"sort"`
		Date1         string   `json:"date1"`
		Date2         string   `json:"date2"`
		Limit         int      `json:"limit"`
		Offset        int      `json:"offset"`
		Group         string   `json:"group"`
		AutoGroupSize string   `json:"auto_group_size"`
		AttrName      string   `json:"attr_name"`
		Quantile      string   `json:"quantile"`
		OfflineWindow string   `json:"offline_window"`
		Attribution   string   `json:"attribution"`
		Currency      string   `json:"currency"`
		AdfoxEventID  string   `json:"adfox_event_id"`
	} `json:"query"`
	Data []struct {
		Dimensions []struct {
			Name    string `json:"name"`
			Favicon string `json:"favicon"`
			URL     string `json:"url,omitempty"`
			ID      string `json:"id,omitempty"`
		} `json:"dimensions"`
		Metrics []float64 `json:"metrics"`
	} `json:"data"`
	TotalRows             int       `json:"total_rows"`
	TotalRowsRounded      bool      `json:"total_rows_rounded"`
	Sampled               bool      `json:"sampled"`
	ContainsSensitiveData bool      `json:"contains_sensitive_data"`
	SampleShare           float64   `json:"sample_share"`
	SampleSize            int       `json:"sample_size"`
	SampleSpace           int       `json:"sample_space"`
	DataLag               int       `json:"data_lag"`
	Totals                []float64 `json:"totals"`
	Min                   []float64 `json:"min"`
	Max                   []float64 `json:"max"`
}

func pars_serchPhras(body []byte) {
	var v AutoGenerated2
	json.Unmarshal(body, &v)
	out := 0.0
	for _, v2 := range v.Data {
		//fmt.Println(v.Data[v1])
		//fmt.Println(v2.Metrics[1])
		out += v2.Metrics[1]

	}
	H2 = serchPhras{out}
	//return //out
}
