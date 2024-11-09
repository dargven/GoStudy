package save

import (
	// для краткости даем короткий алиас пакету
	resp "GoStudy/internal/lib/api/response"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}
type URLSaver interface {
	SaveURL(URL, alias string) (int64, error)
}
