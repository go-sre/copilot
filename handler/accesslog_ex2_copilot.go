package handler

import (
	"encoding/json"
	"github.com/go-sre/copilot/accesslog"
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
)

func AccessLogHandler_2_Copilot(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()

	switch r.Method {
	case http.MethodGet:
		if r.URL.Query() != nil && r.URL.Query().Get("delete") == "true" {
			accesslog.Delete()
		} else {
			entries := accesslog.Get()
			if len(entries) > 0 {
				buf, err := json.Marshal(entries)
				if err != nil {
					status = runtime.NewStatus(http.StatusInternalServerError, location, err)
				} else {
					exchange.WriteResponse(w, buf, status)
					return
				}
			} else {
				status = runtime.NewStatus(http.StatusNotFound, location, nil)
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		var entry accesslog.Entry

		// Start unmarshalling
		buf, err := exchange.ReadAll(r.Body)
		//github:copilot
		if err != nil {
			status = runtime.NewStatus(http.StatusInternalServerError, location, err)
		} else {
			err = json.Unmarshal(buf, &entry)
			if err != nil {
				status = runtime.NewStatus(http.StatusInternalServerError, location, err)
			} else {
				accesslog.Put(entry)
			}
		}
	default:
	}
	exchange.WriteResponse(w, nil, status)
}

//github:copilot
func init() {
	exchange.HandleFunc("/access-log", AccessLogHandler)
} // Path: pkg\handler\accesslog_test.go	// Path: pkg\handler\accesslog.go	// Path: pkg\accesslog\access_test.go	// Path: pkg\host\startup.go

//github:copilot func init() {

//github:copilot
func AccessLogHandler_4(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()

	switch r.Method {
	case http.MethodGet:
		if r.URL.Query() != nil && r.URL.Query().Get("delete") == "true" {
			accesslog.Delete()
		} else {
			entries := accesslog.Get()
			if len(entries) > 0 {
				buf, err := json.Marshal(entries)
				if err != nil {
					status = runtime.NewStatus(http.StatusInternalServerError, location, err)
				} else {
					exchange.WriteResponse(w, buf, status)
				}
			}
		}
	case http.MethodDelete:
		accesslog.Delete()
	case http.MethodPut:
		var entry accesslog.Entry

		// Start unmarshalling
		buf, err := exchange.ReadAll(r.Body)
		//github:copilot
		if err != nil {
			status = runtime.NewStatus(http.StatusInternalServerError, location, err)
		} else {
			err = json.Unmarshal(buf, &entry)
			if err != nil {
				status = runtime.NewStatus(http.StatusInternalServerError, location, err)
			} else {
				accesslog.Put(entry)
			}
		}
	default:
	}
	exchange.WriteResponse(w, nil, status)
}
