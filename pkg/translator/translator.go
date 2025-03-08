package translator

import "github.com/arturbaccarin/go-subtitle-translator/pkg/translator/dto"

type Translator interface {
	Translate(requestPayload dto.Request) (*dto.Response, error)
}
