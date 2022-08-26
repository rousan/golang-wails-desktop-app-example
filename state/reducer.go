package state

import (
	"fmt"
)

func reducer(state *AppState, ac Action) (newState *AppState, err error) {
	defer func() {
		if e := recover(); e != nil {
			newState = nil
			err = fmt.Errorf("Error: %v", e)
		}
	}()

	actionType, payload := ac.Type, ac.Payload

	switch actionType {
	case actionTypeNewDownload:
		return addNewDownload(state, payload.(*ActionPayloadNewDownload))
	case actionTypeProgressUpdate:
		return updateProgress(state, payload.(*ActionPayloadProgressUpdate))
	default:
		return state, nil
	}
}

func addNewDownload(state *AppState, info *ActionPayloadNewDownload) (newState *AppState, err error) {
	newState = state.Clone()
	newState.Downloads = append(newState.Downloads, *info.DownloadInfo)

	return
}

func updateProgress(state *AppState, pUpdate *ActionPayloadProgressUpdate) (newState *AppState, err error) {
	newState = state.Clone()
	for idx := range newState.Downloads {
		curDownloadInfo := &newState.Downloads[idx]

		if curDownloadInfo.ID == pUpdate.DownloadID {
			curDownloadInfo.Status = pUpdate.Status

			if pUpdate.Msg != "" {
				curDownloadInfo.Msg = pUpdate.Msg
			}

			if pUpdate.ContentLength != nil {
				curDownloadInfo.ContentLength = pUpdate.ContentLength
			}

			curDownloadInfo.Progress = pUpdate.Progress
			curDownloadInfo.SavedFilePath = pUpdate.SavedFilePath

			break
		}
	}

	return
}
