package main

import (
	"changeme/state"
	"changeme/utils"
	"context"
	"log"
)

type binds struct {
	ctx context.Context
}

func newBinds() *binds {
	return &binds{
		ctx: nil,
	}
}

func (b *binds) setContext(ctx context.Context) {
	b.ctx = ctx
}

func (b *binds) RequestDownloadURL(urlStr string) string {
	downloadId := utils.GenUUID()

	go func() {
		ac := state.ActionNewDownload(&state.ActionPayloadNewDownload{
			DownloadInfo: &state.DownloadInfo{
				ID:            downloadId,
				URL:           urlStr,
				Status:        state.DownloadStateLoading,
				Msg:           "",
				ContentLength: nil,
				Progress:      0,
				SavedFilePath: "",
			},
		})
		err := state.Dispatch(ac)
		if err != nil {
			log.Printf("failed to dispatch an action: action: %v, error: %v", ac.Type, err)
			return
		}

		savedFilePath, err := downloadURL(
			b.ctx,
			urlStr,
			func(dp downloadProgress) {
				ac := state.ActionProgressUpdate(&state.ActionPayloadProgressUpdate{
					DownloadID:    downloadId,
					Status:        state.DownloadStateLoading,
					Msg:           "",
					ContentLength: dp.contentLength,
					Progress:      dp.progress,
					SavedFilePath: "",
				})
				err := state.Dispatch(ac)
				if err != nil {
					log.Printf("failed to dispatch an action: action: %v, error: %v", ac.Type, err)
					return
				}
			},
		)

		if err != nil {
			ac := state.ActionProgressUpdate(&state.ActionPayloadProgressUpdate{
				DownloadID: downloadId,
				Status:     state.DownloadStateFailed,
				Msg:        err.Error(),
			})
			err := state.Dispatch(ac)
			if err != nil {
				log.Printf("failed to dispatch an action: action: %v, error: %v", ac.Type, err)
				return
			}
		} else {
			ac := state.ActionProgressUpdate(&state.ActionPayloadProgressUpdate{
				DownloadID:    downloadId,
				Status:        state.DownloadStateDone,
				SavedFilePath: savedFilePath,
			})
			err := state.Dispatch(ac)
			if err != nil {
				log.Printf("failed to dispatch an action: action: %v, error: %v", ac.Type, err)
				return
			}
		}
	}()

	return downloadId
}
