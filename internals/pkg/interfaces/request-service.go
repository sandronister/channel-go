package interfaces

import (
	"net/http"
)

type RequestService interface {
	Get(url string) (*http.Request, error)
}
