import { BaseClusterWorker } from 'eris-fleet';
import { Message } from 'eris';
import * as Sentry from '@sentry/node';
import { Setup } from 'eris-fleet/dist/clusters/BaseClusterWorker';
import { readdirSync } from 'fs';
import { join } from 'path';
import { inspect } from 'util';
import Collection from '@discordjs/collection';

import {
  cacheStats,
  handleError,
  hasSecurityClearance,
  logger,
  sendMessage,
} from './utils';
import { Command, properUsage } from './structures';
import { CommandObjects, Commands } from './@types/command';
import PromClient from 'prom-client';
// import { HandleMessageObject } from './@types/handleMessage';

export default class JanusWorker
  extends BaseClusterWorker
  implements Commands {
  Commands = new Collection<string, Command>();

  constructor(setup: Setup) {
    // Worker Cluster.

    // Do not delete this super.
    super(setup);

    /**
     * Set sentry tags
     */
    // set worker id
    Sentry.setTag('workerID', this.workerID);

    logger.debug(this.workerID);
    // set cluster id
    Sentry.setTag('clusterID', this.clusterID);

    //  launch bot
    this.launch();
  }

  /**
   * Launches the necessary services for the bot
   */
  async launch(): Promise<void> {
    await this.loadcmds();

    // Start to collect metrics for Prometheus
    PromClient.collectDefaultMetrics({
      prefix: 'janus_',
      labels: {
        NODE_APP_INSTANCE: process.env.NODE_APP_INSTANCE,
      },
    });

    await cacheStats();

    this.bot.on('messageCreate', (message) =>
      this.handleMessage(message),
    );
  }

  /**
   * Handles inbound messages
   * @param {Message} msg
   */
  async handleMessage(msg: Message): Promise<void> {
    // if bot or self ignore message
    if (
      msg.author.bot ||
      msg.author.id === this.bot.user.id
    )
      return;

    const prefixNormal = process.env.PREFIX;
    const prefixMention = new RegExp(
      `^<@!${this.bot.user.id}>`,
    );

    // mentioning bot
    if (msg.content.match(prefixMention)) {
      // handle bot being mentioned
      await this.botMentioned(msg);
    } else {
      //  normal message handling

      // find matching prefix
      // const prefix = data.msg.content.match(prefixMention)
      //   ? // @ts-ignore
      //   data.msg.content.match(prefixMention)[0]
      //   : prefixNormal;

      const prefix = prefixNormal;

      // Parse args
      const args = msg.content
        .slice(prefix.length)
        .trim()
        .split(/ +/);
      // get command name and make lowercase
      const commandName = args.shift()?.toLowerCase();

      logger.debug(`Command name: ${commandName}`);
      logger.debug(`Args: ${inspect(args)}`);

      if (commandName === undefined) return;

      // const command = this.Commands.has(commandName) ?
      //   this.Commands.get(commandName) :
      //   this.Commands.find(
      //     (cmd) =>
      //       cmd.aliases !== undefined &&
      //       cmd.aliases.includes(commandName),
      //   );

      const command =
        this.Commands.get(commandName) ||
        this.Commands.find(
          (cmd) =>
            cmd.aliases !== undefined &&
            cmd.aliases.includes(commandName),
        );

      if (!command) return undefined;

      // log command was found
      logger.debug(`Found command: ${command.name}`);

      // Check and error if args required but no provided
      // logger.info(`${args.length}`);
      if (command.args && !args.length) {
        await sendMessage(
          this.bot,
          msg.channel.id,
          properUsage(command.name, command.usage),
        );
        return;
      }

      // start transaction for cmd
      const commandTransaction = Sentry.startTransaction({
        name: 'Start Command',
        op: 'command',
      });

      // set user for cmd
      Sentry.setUser({
        id: msg.author.id,
        guildID: msg.guildID,
      });

      // tell sentry what cmd we are running and with what args
      Sentry.addBreadcrumb({
        category: 'command',
        message: `Running command ${command.name}`,
        level: Sentry.Severity.Info,
        data: {
          args: args,
        },
      });

      // if cmd requires security clearance
      if (command.securityClearance !== undefined) {
        // if user has required perms
        if (
          await hasSecurityClearance(
            msg.author.id,
            command.securityClearance,
          )
        ) {
          // tell sentry user has perms for authenticated cmd
          Sentry.addBreadcrumb({
            category: 'command',
            message: `User has perms for ${command.name}`,
            level: Sentry.Severity.Info,
          });

          logger.info(
            `User ${msg.author.id} has perms for ${command.name}`,
          );
        } else {
          // tell sentry someone tried to run an authenticated cmd
          Sentry.addBreadcrumb({
            category: 'command',
            message: `User without proper perms tried to run ${command.name}`,
            level: Sentry.Severity.Log,
          });

          logger.info(
            `User ${msg.author.id} lacks perms for ${command.name}`,
          );

          // end cmd transaction because cmd will not be run
          commandTransaction.finish();

          // finish this handling
          return;
        }
      }

      // const perms = command.channelPermissions();
      // check perms for cmd here
      //
      //
      //
      //
      //

      logger.debug(`Running command ${command.name}`);

      try {
        // execute command
        await command.execute(msg, args);
      } catch (error) {
        handleError(error, 'Command Error');
      } finally {
        // finish transaction
        commandTransaction.finish();
      }

      // clear user from scope
      Sentry.configureScope((scope) => scope.setUser(null));
    }

    // const reply: MessageContent = <MessageContent>(
    //   await this.ipc.command('HandleMessage', data, true)
    // );
    // await this.bot.createMessage(msg.channel.id, reply);
  }

  private async botMentioned(msg: Message): Promise<void> {
    const commandName = 'help';
    const command = this.Commands.get(commandName);

    if (!command) return;

    await command.execute(msg, []);
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

        try {
          const commandObjects: CommandObjects = {
            bot: this.bot,
            ipc: this.ipc,
            workerID: this.workerID,
            clusterID: this.clusterID,
          };

          // import is technically the type 'CommandImport', but is negated through #default
          const command: Command = new (
            await import(`./commands/${folder}/${file}`)
          ).default(commandObjects);

          // logger.debug(
          //   inspect(
          //     (await import(`./commands/${folder}/${file}`))
          //       .default,
          //     false,
          //     6,
          //     true,
          //   ),
          // );

          // log cmd object
          // logger.debug(
          //   `Command: ${inspect(command, false, 2, true)}`,
          // );

          // if command isn't disabled
          if (!command.disabled) {
            // register command
            this.Commands.set(command.name, command);

            // log cmd was loaded
            logger.debug(
              `Loaded the command ${command.name} from module ${folder}`,
            );
          }
        } catch (error) {
          handleError(error);
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
    // Shuts down the instance of the bot

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
