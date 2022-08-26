import React from 'react';
import { Progress } from 'antd';
import styles from './index.module.css';
import { DownloadInfo } from '../../store/downloadsSlice';

interface DownloadItemProps {
  downloadInfo: DownloadInfo;
}

export default function DownloadItem({ downloadInfo }: DownloadItemProps) {
  let status: "success" | "exception" | "active";
  let percent: number;
  let isConnecting = false;

  switch (downloadInfo.status) {
    case "DONE":
      status = "success";
      percent = 100;
      break;
    case "FAILED":
      status = "exception";
      percent = 100;
      break;
    case "LOADING":
      status = "active";
      percent = +(downloadInfo.contentLength ? (downloadInfo.progress / downloadInfo.contentLength) * 100 : 0).toFixed(2);
      isConnecting = !downloadInfo.contentLength;
      break;
  }

  return (
    <div className={styles.downloadItem}>
      <div>
        <Progress
          strokeColor={{
            from: '#108ee9',
            to: '#87d068',
          }}
          percent={percent}
          status={status}
        />
      </div>
      <div>
        {
          isConnecting ? "Connecting..." : ""
        }
      </div>
      <div className={styles.errorMsg}>
        {
          downloadInfo.msg
        }
      </div>
    </div>
  );
}