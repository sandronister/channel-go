package interfaces

type HttpService interface {
	Do(url string) ([]byte, error)
}
