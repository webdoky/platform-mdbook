import path from 'path';

import { kuma } from '@webdoky/yari-ports';

import Preprocessor from '../preprocessor/preprocessor.js';
import warn from '../helpers/warn.js';
import failedMacro from './failed-macro.js';
import expungedMacro from './expunged-macro.js';
import getFrontmatterData from '../helpers/get-frontmatter-data.js';

const { macros: Macros, parseMacroArgs, extractMacros } = kuma;

const DUMMY_REGISTRY = { getPageBySlug: () => ({ tags: [] }) };
const NAVIGATIONAL_MACROS = [
  'cssref',
  'jssidebar',
  'jsref',
  'htmlref',
  'svgref',
  'glossarysidebar',
  'htmlsidebar',
  'listsubpagesforsidebar',
];

/**
 *
 * @param {import('../preprocessor/preprocessor').Section} section
 */
async function applyMacrosToSection(section, registry, srcFolder) {
  const recognizedMacros = extractMacros(section.Chapter.content);
  const frontmatterData = await getFrontmatterData(
    path.resolve(srcFolder, section.Chapter.source_path),
  );
  //   warn(JSON.stringify(frontmatterData, null, 2));
  const macrosRegistry = new Macros({
    env: {
      ...frontmatterData,
      browserCompat: frontmatterData['browser-compat'],
      'browser-compat': undefined,
      path: section.Chapter.source_path,
      targetLocale: 'uk',
      title: section.Chapter.name,
    },
    registry,
  });
  recognizedMacros
    .map((expression) => {
      const { match, functionName, args } = expression;
      let result = match;
      if (NAVIGATIONAL_MACROS.includes(functionName.toLowerCase())) {
        return [match, ''];
      }
      const macroFunction = macrosRegistry.lookup(functionName);
      if (!macroFunction) {
        warn(`Unknown macro: ${match}`);
        return [match, expungedMacro(match)];
      }
      try {
        if (args) {
          result = macroFunction(...parseMacroArgs(args));
        } else {
          result = macroFunction();
        }
      } catch (error) {
        warn(`${error}`);
        warn(`Failed macro: ${match}`);
        return [match, failedMacro(match)];
      }
      return [match, result];
    })
    .forEach((result) => {
      if (!result) {
        return;
      }
      const [match, replacement] = result;
      // eslint-disable-next-line no-param-reassign
      section.Chapter.content = section.Chapter.content.replaceAll(
        match,
        replacement,
      );
    });
  await Promise.all(
    section.Chapter.sub_items.map((subItem) => {
      if (typeof subItem === 'string') {
        return null;
      }
      return applyMacrosToSection(subItem, registry, srcFolder);
    }),
  );
}

/**
 *
 * @param {import('../preprocessor/preprocessor').Book} book
 * @param {import('../preprocessor/preprocessor').Context} context
 */
async function applyMacros(book, context) {
  const srcFolder = context.config.book.src;
  for (const section of book.sections) {
    if (typeof section === 'string') {
      continue;
    }
    await applyMacrosToSection(section, DUMMY_REGISTRY, srcFolder);
  }
  return book;
}

const preprocessor = new Preprocessor(applyMacros);

preprocessor.run();
