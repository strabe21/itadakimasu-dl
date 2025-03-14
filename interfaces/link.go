package interfaces

import (
	"itadakimasu-dl/config"
)

type ILink interface {
	GetUrl() string
	GetServerName() string
	Download(outputFile string) error
	GetId() string
	GetConfig() *config.Link
}
