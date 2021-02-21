// import { BaseServiceWorker } from 'eris-fleet';
// import { Setup } from 'eris-fleet/dist/services/BaseServiceWorker';
// import * as Sentry from '@sentry/node';
// import { HandleMessageObject } from '../@types/handleMessage';
// import { Message, MessageContent } from 'eris';
// import { CommandsCollection } from '../@types/command';
// import { logger } from '../utils';
// import { inspect } from 'util';
//
// export class ServiceWorker extends BaseServiceWorker {
//   constructor(setup: Setup) {
//     // Do not delete this super.
//     super(setup);
//
//     /**
//      * Set sentry tags
//      */
//     // set worker id
//     Sentry.setTag('workerID', this.workerID);
//     // set service name
//     Sentry.setTag('serviceName', this.serviceName);
//
//     // Run this function when your service is ready for use. This MUST be run for the worker spawning to continue.
//     this.serviceReady();
//   }
//
//   // This is the function which will handle commands
//   // eslint-disable-next-line @typescript-eslint/ban-ts-comment
//   // @ts-ignore
//   async handleCommand(
//     data: HandleMessageObject,
//   ): Promise<MessageContent | undefined> {}
//
//   // eslint-disable-next-line @typescript-eslint/ban-ts-comment
//   // @ts-ignore
//   shutdown(done: Done) {
//     // Optional function to gracefully shutdown things if you need to.
//     done(); // Use this function when you are done gracefully shutting down.
//   }
// }
