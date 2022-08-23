import { Sum } from '../../wailsjs/go/main/binds';

export async function sum(a: number, b: number): Promise<number> {
  return Sum(a, b);
}