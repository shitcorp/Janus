import { Worker } from 'bullmq';
import * as Sentry from '@sentry/node';
import { platform, release, version } from 'os';
import { config } from 'dotenv';

import {
  cacheGameStats,
  logger,
  handleError,
  handleWarning,
} from './utils';

// check if node version is below 14
if (Number(process.version.slice(1).split('.')[0]) < 14) {
  logger.error(
    'Node 14.0.0 or higher is required. Please upgrade Node.js on your computer / server.',
  );
  process.exit(1);
}

config();

// define sentry DSN
// @ts-ignore
const DSN: string | null = process.env.SENTRY_URL;

// if not production
// if (process.env.NODE_ENV !== 'production') {
// (async () => {
//   // set DSN to null
//   // DSN = null;
//   // get env vars from .env file
//   (await import('dotenv')).config();
// })();
// }

// start sentry
Sentry.init({
  // @ts-ignore
  dsn: DSN,
  release: `janus-bull-worker@${process.env.npm_package_version}`,
  environment: process.env.NODE_ENV || 'dev',
  maxBreadcrumbs: 100,
  serverName: process.env.serverName || 'dev',

  beforeSend(event) {
    // if user
    if (event.user) {
      // scrub any possible sensitive data

      // don't send email address
      delete event.user.email;
      // don't send username
      delete event.user.username;
      // don't send ip address
      delete event.user.ip_address;
    }
    return event;
  },

  tracesSampler: (samplingContext) => {
    switch (samplingContext.transactionContext.op) {
      case 'shutdown':
        return 1;
      case 'clean cache':
        return 0.5;
      default:
        return 0.5;
    }
  },
});

// set sentry tags
Sentry.setTag('platform', platform());
Sentry.setTag('os.name', version());
Sentry.setTag('os', version() + ' ' + release());
Sentry.setTag('node', process.version);

// @ts-ignore
const cacheWorker = new Worker('Cache', async (job) => {
  if (job.name === 'Game Stats') {
    const result = await cacheGameStats();

    if (!result) new Error('Failed to cache game stats');
  }
});

logger.info('Ready');

// Capture unhandledRejections
process.on('unhandledRejection', (error: Error) => {
  // If not a permission denied discord api error
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  if (error.code !== 50013) {
    handleError(error, 'Unhandled Rejection');
  }
  // logger.error(error);
});

// Capture uncaughtExceptionMonitors
process.on(
  'uncaughtExceptionMonitor',
  (error: Error, origin: string) => {
    handleError(error, 'Uncaught Exception Monitor');

    // // Log
    // logger.error(
    //   `Uncaught Exception Monitor: ${inspect(
    //     error,
    //   )}\n${origin}`,
    // );
    // // Send to sentry
    // Sentry.captureException(error);
  },
);

// Capture warnings
process.on('warning', (warning: Error) => {
  handleWarning(warning);
});

// Close event sent
process.on('SIGTERM', async () => {
  logger.info('SIGTERM signal received.');
  // close workers
  await cacheWorker.close();

  // close Sentry connection
  await Sentry.close(4000);

  // exit
  process.exit(0);
});
