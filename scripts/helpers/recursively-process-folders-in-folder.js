import { readdir } from 'fs/promises';
import path from 'path';
import warn from './warn.js';

/**
 *
 * @param {string} folderPath
 * @param {(subfolder: string, depth: number) => Promise<void>} processFolder
 */
export default async function recursivelyProcessFoldersInFolder(
  folderPath,
  processFolder,
  depth = 0,
) {
  warn(folderPath);
  const result = await processFolder(folderPath, depth);
  if (!result) {
    return;
  }
  const files = await readdir(folderPath, { withFileTypes: true });
  for (const file of files) {
    if (file.isDirectory()) {
      await recursivelyProcessFoldersInFolder(
        path.resolve(folderPath, file.name),
        processFolder,
        depth + 1,
      );
    }
  }
}
