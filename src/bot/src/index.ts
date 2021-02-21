import { isMaster } from 'cluster';
import { Fleet } from 'eris-fleet';
import { join } from 'path';
import * as Sentry from '@sentry/node';
import { platform, version, release } from 'os';
import Redis from 'ioredis';
import PromClient from 'prom-client';
import {
  handleError,
  handleWarning,
  logger,
} from './utils';
import { config } from 'dotenv';
import http from 'http';
// import * as Tracing from '@sentry/tracing';
// import { Queue, QueueScheduler } from 'bullmq';

// @ts-ignore
// const cacheQueueScheduler = new QueueScheduler(
//   'Clean Caches',
// );
// const cacheQueue = new Queue('Clean Caches');

// load dot env config
config();

// check if node version is below 14
if (Number(process.version.slice(1).split('.')[0]) < 14) {
  logger.error(
    'Node 14.0.0 or higher is required. Please upgrade Node.js on your computer / server.',
  );
  process.exit(1);
}

// define sentry DSN
const DSN: string | null = process.env.SENTRY_URL;
// define redis connection
const redis = new Redis();
const register = new PromClient.AggregatorRegistry();

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
      case 'cacheStats':
        return 0.4;
      case 'command':
        return 0.1;
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
  lessLogging: process.env.NODE_ENV === 'production',
  serviceTimeout: 6000,
  clientOptions: {
    allowedMentions: {
      everyone: false,
    },
  },
  startingStatus: {
    status: 'online',
    game: {
      name: '*help | ',
      type: 0,
    },
  },
  // services: [
  //   {
  //     name: 'HandleMessage',
  //     path: join(__dirname, './services/HandleMessage.js'),
  //   },
  // ],
});

// Code to only run on master process
if (isMaster) {
  // standard logs
  Admiral.on('log', (m) => logger.info(m));
  // debug info
  // Admiral.on('debug', (m) => logger.debug(m));
  // warnings
  Admiral.on('warn', (m) => {
    handleWarning(m);
  });
  // errors
  Admiral.on('error', (m) => {
    handleError(m);
  });

  // Logs stats when they arrive
  Admiral.on('stats', async (m) => {
    const pipeline = redis.pipeline();

    // set values
    pipeline
      .set('guilds', m.guilds)
      .set('users', m.users)
      .set('clustersRam', m.clustersRam)
      .set('servicesRam', m.servicesRam)
      .set('masterRam', m.masterRam)
      .set('totalRam', m.totalRam)
      .set('voice', m.voice)
      .set('largeGuilds', m.largeGuilds)
      .set('shardCount', m.shardCount)
      .set('clusters', m.clusters)
      .set('services', m.services);

    // execute pipeline
    pipeline.exec().catch((err) => {
      handleError(err);
    });

    // logger.info(`stats ${inspect(m)}`);
  });

  // (async () => {
  //   // Repeat job once every day at 3:15 (am)
  //   await cacheQueue.add('clean', undefined, {
  //     repeat: {
  //       cron: '15 3 * * *',
  //     },
  //   });
  // })();

  http
    .createServer(async (req, res) => {
      //handle every single request with this callback

      try {
        const metrics = await register.clusterMetrics();

        // send content type
        res.writeHead(200, {
          'Content-Type': register.contentType,
        });
        // send metrics
        res.write(metrics);
      } catch (error) {
        handleError(error);

        // set internal error status code
        res.statusCode = 500;
        // send error message
        res.write(error.message);
      } finally {
        // end the response
        res.end();
      }
    })
    .listen(process.env.PORT, () =>
      logger.info(
        `Prometheus running at http://localhost:${process.env.PORT}`,
      ),
    );
}

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
