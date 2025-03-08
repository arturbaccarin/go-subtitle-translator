package dto

type Request struct {
	Text               []string `json:"text"`
	TargetLang         string   `json:"target_lang"`
	SourceLang         string   `json:"source_lang,omitempty"`
	ModelType          string   `json:"model_type,omitempty"`
	PreserveFormatting bool     `json:"preserve_formatting,omitempty"`
}

type Response struct {
	Translations []TranslationResult `json:"translations"`
}

type TranslationResult struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}
