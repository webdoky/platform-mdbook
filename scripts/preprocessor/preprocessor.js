import process from 'process';

import readStdinUntilEof from './read-stdin-until-eof.js';
import writeLargeDataToStdout from './write-large-data-to-stdout.js';

/**
 * @typedef {Object} Chapter
 * @property {string} content
 * @property {string} name
 * @property {string} path
 * @property {string} source_path
 * @property {(string | Section)[]} sub_items
 */

/**
 * @typedef {Object} Section
 * @property {Chapter} Chapter
 *
 */

/**
 * @typedef {Object} Book
 * @property {(Section | "Separator")[]} sections
 */

/**
 * @typedef {Object} Context
 * @property {Object} config
 * @property {Object} config.book
 * @property {string} config.book.src
 * @property {string} root
 */

export default class Preprocessor {
  /**
   *
   * @param {(book: Book, context: Context) => Promise<Book>} processBook
   */
  constructor(processBook) {
    this.process = processBook;
  }

  /**
   * @returns {Promise<void>}
   */
  async run() {
    try {
      if (process.argv[2] === 'supports' && process.argv[3] === 'html') {
        process.exit(0);
      }
      const jsonInput = await readStdinUntilEof();
      const [context, book] = JSON.parse(jsonInput);
      const processedBook = await this.process(book, context);
      await writeLargeDataToStdout(JSON.stringify(processedBook));
      process.exit(0);
    } catch (error) {
      process.stderr.write(`${error}\n`, () => {
        process.exit(1);
      });
    }
  }
}
