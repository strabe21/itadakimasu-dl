package ui

import (
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/text"
)

var GetProgressWriter progress.Writer = loadProgressWriter()

func loadProgressWriter() progress.Writer {
	pw := progress.NewWriter()
	pw.SetAutoStop(true)
	pw.SetTrackerLength(50)
	pw.SetMessageLength(50)
	pw.SetStyle(progress.StyleBlocks)
	pw.SetTrackerPosition(progress.PositionLeft)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Visibility.Time = false
	pw.Style().Visibility.Speed = true
	pw.Style().Colors = progress.StyleColors{
		Message: text.Colors{text.FgHiBlack},
		Percent: text.Colors{text.FgHiBlue},
		Tracker: text.Colors{text.FgHiWhite},
		Speed:   text.Colors{text.FgHiGreen},
		Value:   text.Colors{text.FgHiWhite},
		Error:   text.Colors{text.FgHiRed},
	}

	return pw
}

type ProgressWriter struct {
	Tracker *progress.Tracker
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Tracker.Increment(int64(n))
	return n, nil
}
