import { stat } from 'fs/promises';

/**
 *
 * @param {import('fs').PathLike} filePath
 * @returns {Promise<boolean>}
 */
export default async function doesFileExist(filePath) {
  try {
    await stat(filePath);
    return true;
  } catch (error) {
    if (error.code === 'ENOENT') {
      return false;
    }
    throw error;
  }
}
