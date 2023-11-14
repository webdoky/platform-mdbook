import { readFile } from 'fs/promises';

import matter from 'gray-matter';

export default async function getFrontmatterData(markdownFilePath) {
  const markdown = await readFile(markdownFilePath, 'utf8');
  const { data } = matter(markdown);
  return data;
}
