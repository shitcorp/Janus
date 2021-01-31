import { isMaster } from 'cluster';
import { Fleet } from 'eris-fleet';
import { join } from 'path';
import { inspect } from 'util';
import * as Sentry from '@sentry/node';
// import * as Tracing from '@sentry/tracing';
import { platform, version, release } from 'os';

import { config } from 'dotenv';
config();

import { logger } from './utils';

// check if node version is below 14
if (Number(process.version.slice(1).split('.')[0]) < 14) {
  logger.error(
    'Node 14.0.0 or higher is required. Please upgrade Node.js on your computer / server.',
  );
  process.exit(1);
}

// define sentry DSN
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
  dsn: DSN,
  release: `janus-bot@${process.env.npm_package_version}`,
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
      case 'loadCMD':
        return 0.2;
      case 'loadModule':
        return 0.3;
      case 'startManager':
        return 0.5;
      case 'command':
        return 0.05;
      case 'reloadCMD':
        return 0.8;
      case 'shutdown':
        return 1;
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

const Admiral = new Fleet({
  path: join(__dirname, './bot.js'),
  token: process.env.DISCORD_TOKEN,
});

// Code to only run on master process
if (isMaster) {
  // standard logs
  Admiral.on('log', (m) => logger.info(m));
  // debug info
  // Admiral.on('debug', (m) => logger.debug(m));
  // warnings
  Admiral.on('warn', (m) => {
    logger.warn(m);
  });
  // errors
  Admiral.on('error', (m) => {
    logger.error(inspect(m));
  });

  // Logs stats when they arrive
  // Admiral.on('stats', (m) => logger.info(inspect(m)));
}

// Capture unhandledRejections
process.on('unhandledRejection', (error: Error) => {
  // If not a permission denied discord api error
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  if (error.code !== 50013) {
    // Log
    logger.error(`Unhandled Rejection: ${inspect(error)}`);
    // Send to sentry
    Sentry.captureException(error);
  }
  // logger.error(error);
});

// Capture uncaughtExceptionMonitors
process.on(
  'uncaughtExceptionMonitor',
  (error: Error, origin: string) => {
    // Log
    logger.error(
      `Uncaught Exception Monitor: ${inspect(
        error,
      )}\n${origin}`,
    );
    // Send to sentry
    Sentry.captureException(error);
  },
);

// Capture warnings
process.on('warning', (warning: Error) => {
  // Log
  logger.warn(`Warning: ${inspect(warning)}`);
  // Send to sentry
  Sentry.captureException(warning);
});
