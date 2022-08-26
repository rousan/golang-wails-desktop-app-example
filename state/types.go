package state

import "github.com/samber/lo"

const (
	DownloadStateDone    DownloadStatus = "DONE"
	DownloadStateLoading DownloadStatus = "LOADING"
	DownloadStateFailed  DownloadStatus = "FAILED"
)

type AppState struct {
	Downloads []DownloadInfo `json:"downloads"`
}

func (a *AppState) Clone() *AppState {
	newState := new(AppState)

	newState.Downloads = lo.Map(a.Downloads, func(s DownloadInfo, _ int) DownloadInfo {
		return *s.Clone()
	})

	return newState
}

type DownloadInfo struct {
	ID            string         `json:"id"`
	URL           string         `json:"url"`
	Status        DownloadStatus `json:"status"`
	Msg           string         `json:"msg"`
	ContentLength *uint64        `json:"contentLength"`
	Progress      uint64         `json:"progress"`
	SavedFilePath string         `json:"savedFilePath"`
}

func (d *DownloadInfo) Clone() *DownloadInfo {
	newInfo := new(DownloadInfo)

	newInfo.ID = d.ID
	newInfo.URL = d.URL
	newInfo.Status = d.Status
	newInfo.Msg = d.Msg
	newInfo.ContentLength = d.ContentLength
	newInfo.Progress = d.Progress
	newInfo.SavedFilePath = d.SavedFilePath

	return newInfo
}

type DownloadStatus string

type subscriberCallback = func(newState *AppState) error

type subscriberInfo struct {
	id       string
	callback subscriberCallback
}
