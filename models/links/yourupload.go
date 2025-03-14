package links

import (
	"fmt"
	"itadakimasu-dl/config"
	"itadakimasu-dl/interfaces"
	"itadakimasu-dl/network"
	"math/rand/v2"
	"time"
)

type yourupload struct {
	Url  string
	Id   string
	Name string
	Conf *config.Link
}

func (y *yourupload) Download(outputFile string) error {

	time.Sleep(time.Duration(rand.IntN(1000)))
	doc, err := network.HttpGetDocument(y.Url)
	if err != nil {
		return err
	}
	videoURL, exists := doc.Find(`meta[property="og:video"]`).Attr("content")
	if !exists {
		return fmt.Errorf("video URL not found from %s", y.Url)
	}
	return network.DownloadFile(videoURL, outputFile, "https://yourupload.com/")

}

func (y *yourupload) GetConfig() *config.Link {
	return y.Conf

}

func (y *yourupload) GetId() string {
	return y.Id
}

func (y *yourupload) GetServerName() string {
	return y.Name
}

func (y *yourupload) GetUrl() string {
	return y.Url
}

func NewYourUpload(url, name string, conf *config.Link) interfaces.ILink {
	return &yourupload{
		Url:  url,
		Name: name,
		Conf: conf,
	}
}
