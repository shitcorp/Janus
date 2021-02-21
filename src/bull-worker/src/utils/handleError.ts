import * as Sentry from '@sentry/node';
import { logger } from '../utils';
import { inspect } from 'util';

/**
 * Handles basic error logging and capturing
 * @param error
 * @param logPrefixMessage
 */
export const handleError = (
  error: any,
  logPrefixMessage?: string | undefined,
  // ...args: string[]
) => {
  // send error to sentry
  Sentry.captureException(error);

  const inspected = inspect(error, false, null, true);

  // let compileArgs = '';
  // args.forEach((arg: any) => {
  //   compileArgs += arg;
  //   compileArgs += "\n"
  // });

  if (logPrefixMessage) {
    // log error
    logger.error(`${logPrefixMessage}: ${inspected}`);
  } else {
    logger.error(inspected);
  }
};
