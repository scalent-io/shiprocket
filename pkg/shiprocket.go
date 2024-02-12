package pkg

type ShiprockertService interface {
	GetToken(baseURL, email, password string) (string, error)
}
