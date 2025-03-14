package interfaces

type IAnime interface {
	GetEpisodeByNumber(number int) (IEpisode, error)
	SetEpisodesByList(list []int) error
	FetchEpisodes() error
	GetId() int
	GetEpisodes() map[int]IEpisode
	GetName() string
	GetUrl() string
	Download(animePath string)
	GetSource() string
}
