import { BaseClusterWorker } from 'eris-fleet';
import { Message } from 'eris';
import * as Sentry from '@sentry/node';
import { Setup } from 'eris-fleet/dist/clusters/BaseClusterWorker';
// import { promisify } from 'util';
import { readdirSync } from 'fs';
// import { readdirSync } from 'fs';
import { join } from 'path';
import { inspect } from 'util';

import { logger } from './utils';
import { Command } from './structures';

export default class JanusWorker extends BaseClusterWorker {
  private Commands = new Map<string, Command>();

  constructor(setup: Setup) {
    // Worker Cluster.

    // Do not delete this super.
    super(setup);

    // set sentry tags
    Sentry.setTag('workerID', this.workerID);
    Sentry.setTag('clusterID', this.clusterID);

    //  launch bot
    this.launch();
  }

  /**
   * Launches the necessary services for the bot
   */
  async launch(): Promise<void> {
    await this.loadcmds();

    this.bot.on('messageCreate', (message) =>
      this.handleMessage(message),
    );
  }

  /**
   * Handles inbound messages
   * @param {Message} msg
   */
  async handleMessage(msg: Message): Promise<void> {
    if (msg.content === '!ping' && !msg.author.bot) {
      await this.bot.createMessage(msg.channel.id, 'Pong!');
    }
  }

  /**
   * Loads/reloads all commands
   */
  async loadcmds() {
    // loads the commands

    // loading cmds breadcrumb
    Sentry.addBreadcrumb({
      category: 'loadingCMDs',
      message: 'Loading Commands',
      level: Sentry.Severity.Info,
    });

    const folders = await readdirSync(
      join(__dirname, '/commands/'),
    );

    logger.debug(`Modules: ${inspect(folders)}`);

    for (const folder of folders) {
      const loadModule = Sentry.startTransaction({
        name: 'Load Module',
        op: 'loadModule',
        data: {
          folder,
        },
      });

      // get every command in module folder
      const commands = readdirSync(
        `./dist/commands/${folder}`,
      ).filter((file) => file.endsWith('.js'));

      logger.debug(`Commands: ${inspect(commands)}`);

      for (const file of commands) {
        const loadCommand = Sentry.startTransaction({
          name: 'Load Command',
          op: 'loadCMD',
          data: {
            file,
          },
        });

        // import is technically the type 'CommandImport', but is negated through #default
        const command: Command = (
          await import(`./commands/${folder}/${file}`)
        ).default;

        try {
          // log cmd object
          logger.debug(
            `Command: ${inspect(command, true, 4, true)}`,
          );

          // register command
          this.Commands.set(command.name, command);

          // log cmd was loaded
          logger.debug(
            `Loaded the command ${command.name} from module ${folder}`,
          );
        } catch (error) {
          // log error
          logger.error(inspect(error));
          // send error to sentry
          Sentry.captureException(error);
        } finally {
          // finish cmd load transaction
          loadCommand.finish();
        }
      }

      // finish module load transaction
      loadModule.finish();
    }
  }

  /**
   * Shuts down the instance of the bot
   * @param {Done} done
   */
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  async shutdown(done: Done): void {
    // shutdown breadcrumb
    Sentry.addBreadcrumb({
      category: 'shutdown',
      message: `Shutting down ${this.workerID}`,
      level: Sentry.Severity.Info,
      data: {
        worker: this.workerID,
        cluster: this.clusterID,
      },
    });

    // force sentry to push all data to server
    await Sentry.flush(4000);

    // When done shutting down
    done();
  }
}
