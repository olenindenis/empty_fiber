package response

import (
	"io"
	"log"
	"net/http"
)

func Body(response *http.Response) string {
	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b)
}
