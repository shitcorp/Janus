import {
  AdvancedMessageContent,
  Message,
  Permission,
} from 'eris';
import { Embed } from './Embed';
import {
  CommandChannelPermsArray,
  CommandObjects,
  CommandOptions,
  CommandPerms,
} from '../@types/command';
import {
  handleError,
  sendMessage as sendMessageUtil,
} from '../utils';

/**
 * Represents a command
 */
export class Command
  implements CommandOptions, CommandObjects, CommandPerms {
  readonly bot;
  readonly ipc;
  readonly workerID;
  readonly clusterID;

  readonly name;
  readonly description;
  readonly module;
  readonly aliases?;
  readonly args?;
  readonly usage?;
  readonly disabled?;
  readonly cooldown?;
  readonly securityClearance?;

  readonly channelPermissions: CommandChannelPermsArray;
  readonly userPermissions: string[];
  readonly guildPermissions: string[];

  /**
   * Creates a command
   * @constructor
   * @param commandObjects
   * @param options
   */
  constructor(
    commandObjects: CommandObjects,
    options: CommandOptions,
  ) {
    // save core items for command
    this.bot = commandObjects.bot;
    this.ipc = commandObjects.ipc;
    this.workerID = commandObjects.workerID;
    this.clusterID = commandObjects.clusterID;

    this.name = options.name;
    this.description = options.description;
    this.module = options.module;
    this.aliases = options.aliases;
    this.args = options.args;
    this.usage = options.usage;
    this.disabled = options.disabled;
    this.cooldown = options.cooldown;
    this.securityClearance = options.securityClearance
      ? options.securityClearance
      : undefined;

    this.channelPermissions = options.perms
      ?.channelPermissions
      ? options.perms.channelPermissions
      : ['sendMessages', 'embedLinks'];
    this.userPermissions = options.perms?.userPermissions
      ? options.perms.userPermissions
      : [];
    this.guildPermissions = options.perms?.guildPermissions
      ? options.perms.guildPermissions
      : [];
  }

  /**
   * Executes a given command
   * @param message
   * @param args
   */
  async execute(
    message: Message,
    args: string[],
  ): Promise<void> {
    return;
  }

  /**
   * Checks whether or not the bot has the required channel perms
   * @param permissions eris permissions object
   */
  checkChannelPermissions(
    permissions: Permission,
  ): boolean | string[] {
    const perms = [];

    // loop through all channel perms
    for (const perm of this.channelPermissions) {
      // if doesn't have the permission
      if (!permissions.has(perm)) perms.push(perm);
    }

    if (perms.length > 0) {
      return perms;
    } else {
      return true;
    }
  }

  /**
   * Checks whether or not the bot has the required user perms
   * @param permissions eris permissions object
   */
  checkUserPermissions(permissions: Permission) {
    const perms = [];

    // loop through all user perms
    for (const perm of this.userPermissions) {
      // if doesn't have the permission
      if (!permissions.has(perm)) perms.push(perm);
    }

    if (perms.length > 0) {
      return perms;
    } else {
      return true;
    }
  }

  /**
   * Checks whether or not the bot has the required guild perms
   * @param permissions eris permissions object
   */
  checkGuildPermissions(permissions: Permission) {
    const perms = [];

    // loop through all guild perms
    for (const perm of this.guildPermissions) {
      // if doesn't have the permission
      if (!permissions.has(perm)) perms.push(perm);
    }

    if (perms.length > 0) {
      return perms;
    } else {
      return true;
    }
  }

  /**
   * Handles message sending
   * @param channelID
   * @param message
   */
  async sendMessage(
    channelID: string,
    message: string | Embed | AdvancedMessageContent,
  ): Promise<Message | void> {
    return await sendMessageUtil(
      this.bot,
      channelID,
      message,
    );
  }

  /**
   * Starts the typing status in specified channel
   * @param channelID
   */
  async startTyping(channelID: string) {
    await this.bot
      .sendChannelTyping(channelID)
      .catch((err) => {
        handleError(err);
      });
  }
}
