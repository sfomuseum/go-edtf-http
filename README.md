# go-edtf-http

Go package for sfomuseum/go-edtf HTTP handlers.

## Example

```
import (
	"flag"
	"github.com/sfomuseum/go-edtf-http/api"
	"net/http"
)

func main() {

	api_parse_handler, _ := api.ParseHandler()
	api_valid_handler, _ := api.IsValidHandler()
	api_matches_handler, _ := api.MatchesHandler()

	mux := http.NewServeMux()
	mux.Handle("/api/parse", api_parse_handler)
	mux.Handle("/api/valid", api_valid_handler)
	mux.Handle("/api/matches", api_matches_handler)

	http.ListenAndServe("localhost:8080", mux)
}
```

_Error handling omitted for brevity._

### /api/matches

```
> curl -s 'http://localhost:8080/api/matches?edtf=1985-01-03/1987'
{"level":0,"feature":"Time Interval"}
```

### /api/parse

```
$> curl -s 'http://localhost:8080/api/parse?edtf=1985-01-03/1987' | jq
{
  "start": {
    "edtf": "1985-01-03",
    "lower": {
      "datetime": "1985-01-03T00:00:00Z",
      "timestamp": 473558400,
      "ymd": {
        "year": 1985,
        "month": 1,
        "day": 3
      },
      "precision": 64
    },
    "upper": {
      "datetime": "1985-01-03T23:59:59Z",
      "timestamp": 473644799,
      "ymd": {
        "year": 1985,
        "month": 1,
        "day": 3
      },
      "precision": 64
    }
  },
  "end": {
    "edtf": "1987",
    "lower": {
      "datetime": "1987-01-01T00:00:00Z",
      "timestamp": 536457600,
      "ymd": {
        "year": 1987,
        "month": 1,
        "day": 1
      },
      "precision": 64
    },
    "upper": {
      "datetime": "1987-12-31T23:59:59Z",
      "timestamp": 567993599,
      "ymd": {
        "year": 1987,
        "month": 12,
        "day": 31
      },
      "precision": 64
    }
  },
  "edtf": "1985-01-03/1987",
  "level": 0,
  "feature": "Time Interval"
}
```

### /api/valid

```
$> curl -s 'http://localhost:8080/api/valid?edtf=Jan,%203%201985'
false
```

```
$> curl -s 'http://localhost:8080/api/valid?edtf=1985-01-03'
true
```


## See also

* https://github.com/sfomuseum/go-edtf