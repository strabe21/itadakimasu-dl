package config

import "golang.org/x/sync/semaphore"

type Link struct {
	MaxConcurrentDownloads int                 `json:"maxConcurrentDownloads"`
	Priority               int                 `json:"priority"`
	Semaphore              *semaphore.Weighted `json:"-"`
}
