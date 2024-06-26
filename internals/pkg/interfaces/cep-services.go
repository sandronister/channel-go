package interfaces

type CepServices interface {
	Get(data chan<- string, cep string)
}
