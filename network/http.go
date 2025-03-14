package network

import (
	"fmt"
	"io"
	"itadakimasu-dl/ui"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jedib0t/go-pretty/v6/progress"
)

func HttpGetDocument(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %d", res.StatusCode)
	}
	return goquery.NewDocumentFromReader(res.Body)
}

func DownloadFile(fileURL string, outputFile string, referer string) error {
	client := &http.Client{
		Timeout: 0,
	}
	episodeNameSplitted := strings.Split(outputFile, "/")
	episodeName := episodeNameSplitted[len(episodeNameSplitted)-1]
	tracker := progress.Tracker{
		Message: fmt.Sprintf("Downloading %s", episodeName),
		Total:   0,
		Units:   progress.UnitsBytes,
	}
	ui.GetProgressWriter.AppendTracker(&tracker)
	req, err := http.NewRequest("GET", fileURL, nil)
	if err != nil {
		tracker.MarkAsErrored()
		return fmt.Errorf("failed to create request: %w", err)
	}

	if referer != "" {
		req.Header.Set("Referer", referer)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		tracker.MarkAsErrored()
		return fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		tracker.MarkAsErrored()
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	out, err := os.Create(outputFile)
	if err != nil {
		tracker.MarkAsErrored()
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	tracker.Total = resp.ContentLength
	_, err = io.Copy(out, io.TeeReader(resp.Body, &ui.ProgressWriter{Tracker: &tracker}))
	if err != nil {
		tracker.MarkAsErrored()
		return fmt.Errorf("error writing to output file: %w", err)
	}
	tracker.MarkAsDone()
	return nil
}
