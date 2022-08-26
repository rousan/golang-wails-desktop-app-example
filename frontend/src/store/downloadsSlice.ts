import { createSlice } from '@reduxjs/toolkit';

export interface DownloadInfo {
  id: string;
  url: string;
  status: DownloadState;
  msg: string;
  contentLength: number | null;
  progress: number;
  savedFilePath: string;
}

export type DownloadState = "DONE" | "LOADING" | "FAILED";

const downloadsSlice = createSlice({
  name: 'downloads',
  initialState: [] as DownloadInfo[],
  reducers: {
    syncDownloads: (state, action: { type: string; payload: DownloadInfo[]; }) => {
      const newDownloads = action.payload;
      state.splice(0, state.length);
      state.push(...newDownloads);
    }
  },
});

export const { syncDownloads } = downloadsSlice.actions;

export default downloadsSlice.reducer;