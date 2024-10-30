package xendit

import "net/http"

//go:generate mockgen -source=contract.go -destination=mock/httpconnector_mock.go -package=mock
type HttpConnector interface {
	Do(req *http.Request) (*http.Response, error)
}
