package models

type CrudWeb struct {
	Name      string
	SearchUrl string
	BaseUrl   string
}

func (a *CrudWeb) GetBaseURL() string {
	return a.BaseUrl
}

func (a *CrudWeb) GetSearchUrl() string {
	return a.SearchUrl
}

func (a *CrudWeb) GetName() string {
	return a.Name
}
