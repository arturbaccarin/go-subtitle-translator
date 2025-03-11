package dto

import "github.com/arturbaccarin/go-subtitle-translator/pkg/translator/deepl/language"

type Request struct {
	Text               []string                           `json:"text"`
	TargetLang         language.TranslationTargetLanguage `json:"target_lang"`
	SourceLang         language.TranslationSourceLanguage `json:"source_lang,omitempty"`
	ModelType          string                             `json:"model_type,omitempty"`
	PreserveFormatting bool                               `json:"preserve_formatting,omitempty"`
}

type Response struct {
	Translations []TranslationResult `json:"translations"`
}

type TranslationResult struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}
