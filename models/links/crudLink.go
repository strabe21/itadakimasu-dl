package links

import (
	"fmt"
	"itadakimasu-dl/config"
	"itadakimasu-dl/interfaces"
	"strings"

	"golang.org/x/sync/semaphore"
)

type CrudLink struct {
	Url  string
	Name string
}

func (d *CrudLink) GetConfig() *config.Link {
	return &config.Link{
		MaxConcurrentDownloads: 1,
		Priority:               0,
		Semaphore:              semaphore.NewWeighted(1),
	}
}

func (d *CrudLink) Download(outputFile string) error {

	return fmt.Errorf("server %s not implemented yet", d.Name)
}

func (d *CrudLink) GetId() string {
	return ""
}

func (d *CrudLink) GetServerName() string {
	return d.Name
}

func (d *CrudLink) GetUrl() string {
	return d.Url
}

func NewDefaultLink(url string, name string) interfaces.ILink {
	name = strings.ToLower(name)
	var conf = config.GetConfig.GetLinkConfig(name)
	switch name {
	case "stape":
		return NewStapeLink(url, name, conf)
	case "yourupload":
		return NewYourUpload(url, name, conf)
	default:
		return &CrudLink{
			Url:  url,
			Name: name,
		}
	}
}
