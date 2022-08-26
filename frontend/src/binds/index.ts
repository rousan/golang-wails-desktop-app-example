import { RequestDownloadURL } from '../../wailsjs/go/main/binds';

export async function requestDownloadURL(url: string): Promise<string> {
  return RequestDownloadURL(url);
}