import React, { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { Input } from 'antd';
import { Link } from 'react-router-dom';
import { AppState } from "../../store";
import styles from './index.module.css';
import DownloadItem from "../../components/DownloadItem";
import { requestDownloadURL } from "../../binds";

export default function Home() {
  const [url, setURL] = useState<string>("");
  const downloads = useSelector((state: AppState) => state.downloads);

  const onClickDownload = async () => {
    const id = await requestDownloadURL(url);
    console.log("New download started: ", id);
  };

  return (
    <div className={styles.home}>
      <div className={styles.downloadInputContainer}>
        <Input
          addonBefore="Download URL"
          placeholder="Paste here"
          autoFocus
          value={url}
          onChange={(evt) => setURL(evt.target.value)}
          onPressEnter={onClickDownload}
        />
      </div>
      <div className={styles.downloadsContainer}>
        {
          downloads.map((downloadInfo) => {
            return (
              <DownloadItem
                key={downloadInfo.id}
                downloadInfo={downloadInfo}
              />
            );
          })
        }
      </div>
    </div>
  );
}