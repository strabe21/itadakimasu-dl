package interfaces

import (
	"sync"
)

type IEpisode interface {
	InitDownload(path string, wg *sync.WaitGroup)
	FetchDownloadLinks() error
	GetLinks() []ILink
	GetNumber() int
	GetUrl() string
	GetAnimeName() string
	LinkExist(link ILink) bool
}
