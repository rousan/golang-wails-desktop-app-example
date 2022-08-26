package state

const (
	actionTypeNewDownload    actionType = "NEW_DOWNLOAD"
	actionTypeProgressUpdate actionType = "PROGRESS_UPDATE"
)

type Action struct {
	Type    actionType
	Payload any
}

type actionType string

type ActionPayloadNewDownload struct {
	DownloadInfo *DownloadInfo
}

type ActionPayloadProgressUpdate struct {
	DownloadID    string
	Status        DownloadStatus
	Msg           string
	ContentLength *uint64
	Progress      uint64
	SavedFilePath string
}

func ActionNewDownload(payload *ActionPayloadNewDownload) Action {
	return Action{
		Type:    actionTypeNewDownload,
		Payload: payload,
	}
}

func ActionProgressUpdate(payload *ActionPayloadProgressUpdate) Action {
	return Action{
		Type:    actionTypeProgressUpdate,
		Payload: payload,
	}
}
