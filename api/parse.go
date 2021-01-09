package api

import (
	"net/http"
	"encoding/json"
	"github.com/aaronland/go-http-sanitize"
	"github.com/sfomuseum/go-edtf/parser"
)

func ParseHandler() (http.HandlerFunc, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		edtf_str, err := sanitize.GetString(req, "edtf")

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusBadRequest)
			return
		}

		if edtf_str == "" {
			http.Error(rsp, "Empty '?edtf=' parameter", http.StatusBadRequest)
			return
		}

		edtf_date, err := parser.ParseString(edtf_str)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		
		enc := json.NewEncoder(rsp)
		err = enc.Encode(edtf_date)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil	
}
