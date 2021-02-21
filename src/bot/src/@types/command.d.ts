import { Command } from '../structures';
import Collection from '@discordjs/collection';
import { Client } from 'eris';
import { IPC } from 'eris-fleet/dist/util/IPC';

/**
 * Config items for each command
 */
interface CommandOptions {
  readonly name: string;
  readonly description: string;
  readonly module: string;
  readonly aliases?: Array<string>;
  readonly args?: boolean;
  readonly usage?: string;
  readonly disabled?: boolean;
  readonly cooldown?: 3 | 5 | 7 | 10;
  readonly securityClearance?: 0 | 1 | 2 | 3;

  perms?: CommandPerms;
}

/**
 * A list of required perms for the bot in the channel
 */
interface CommandChannelPermsArray extends Array<string> {
  [index: number]: string;
  0: 'sendMessages';
  1: 'embedLinks';
}

/**
 * Arrays of strings that represent perms
 */
interface CommandPerms {
  readonly channelPermissions: CommandChannelPermsArray;
  readonly userPermissions: string[];
  readonly guildPermissions: string[];
}

/**
 * Required items for commands to work properly
 */
interface CommandObjects {
  bot: Client;
  ipc: IPC;
  clusterID: number;
  workerID: number;
}

/**
 * A discord.js collection that hold commands
 */
type CommandsCollection = Collection<string, Command>;

/**
 * An interface that implements the commands collection
 */
interface Commands {
  Commands: CommandsCollection;
}
