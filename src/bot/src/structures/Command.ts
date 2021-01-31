import { Channel, Client, Message, Permission } from 'eris';
import { Embed } from './Embed';
import { CommandOptions } from '../@types/command';

/**
 * Represents a command
 */
export abstract class Command implements CommandOptions {
  private bot: Client;

  readonly name: string;
  readonly description: string;
  readonly module: string;
  readonly permissions: string[];
  readonly aliases?: Array<string>;
  readonly args?: boolean;
  readonly usage?: string;
  readonly disabled?: boolean;
  readonly cooldown?: 3 | 5 | 7 | 10;

  /**
   * Creates a command
   * @constructor
   * @param bot
   * @param options
   */
  constructor(bot: Client, options: CommandOptions) {
    this.bot = bot;

    this.name = options.name;
    this.description = options.description;
    this.module = options.module;
    this.permissions = options.permissions
      ? [
          'sendMessages',
          'embedLinks',
          ...options.permissions,
        ]
      : ['sendMessages', 'embedLinks'];
    this.aliases = options.aliases;
    this.args = options.args;
    this.usage = options.usage;
    this.disabled = options.disabled;
    this.cooldown = options.cooldown;
  }

  /**
   * Executes a given command
   * @param message
   * @param args
   */
  abstract execute(message: Message, args: string[]): void;

  /**
   * Checks whether or not the passed through permissions are allowed
   * @param permissions eris.js permissions object
   */
  public checkPermissions(
    permissions: Permission,
  ): boolean {
    for (const perm of this.permissions) {
      if (!permissions.has(perm)) {
        return false;
      }
    }
    return true;
  }

  /**
   * Sends message to specified channel
   * @param channel
   * @param message
   * @private
   */
  public async sendMessage(
    channel: Channel,
    message: string | Embed,
  ) {
    if (message instanceof Embed) {
      await this.bot.createMessage(channel.id, {
        embed: { ...message.toJSON() },
      });
    }
  }
}
