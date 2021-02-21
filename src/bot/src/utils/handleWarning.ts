import * as Sentry from '@sentry/node';
import { logger } from '../utils';
import { inspect } from 'util';

/**
 * Handles basic warning logging and capturing
 * @param warning
 * @param logPrefixMessage
 */
export const handleWarning = (
  warning: any,
  logPrefixMessage?: string | undefined,
  // ...args: string[]
) => {
  // send exception to sentry
  Sentry.captureException(warning);

  const inspected = inspect(warning, false, null, true);

  if (logPrefixMessage) {
    // log error
    logger.warn(`${logPrefixMessage}: ${inspected}`);
  } else {
    logger.warn(inspected);
  }
};
