package interfaces

type IWeb interface {
	SearchAnime(query string) ([]IAnime, error)
	GetName() string
	GetSearchUrl() string
	GetBaseURL() string
}
