import { Message } from 'eris';
import {
  Command,
  Embed,
  generalError,
} from '../../structures/';
import { CommandObjects } from '../../@types/command';
import { getUser } from '../../utils';
import { DateTime } from 'luxon';

export default class player extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'player',
      description:
        'Get information about a specific player.',
      module: 'starcitizen',
      aliases: ['user'],
      args: true,
      usage: '<player handle>',
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const user = await getUser(args[0]);

    // If player doesn't exist or data is null
    if (
      user === undefined ||
      user.data === null ||
      Object.keys(user.data).length === 0
    ) {
      await this.sendMessage(
        message.channel.id,
        generalError(
          `No user is under the handle \`${args[0]}\``,
        ),
      );
      return;
    }

    const enlisted = DateTime.fromISO(
      user.data.profile.enlisted,
    );

    const playerEmbed = new Embed()
      .setTitle(user.data.profile.display)
      .setThumbnail(user.data.profile.image)
      .setURL(user.data.profile.page.url)
      .addField('Username:', user.data.profile.handle, true)
      .addField('Title:', user.data.profile.handle, true)
      .addField('Badge:', user.data.profile.badge, true)
      .addField(
        'Enlisted:',
        enlisted.toLocaleString({
          month: 'short',
          day: 'numeric',
          year: 'numeric',
        }),
        true,
      )
      .setTimestamp();

    if (
      Object.keys(user.data.organization).length !== 0 &&
      user.data.organization.name !== undefined
    ) {
      playerEmbed.addField(
        'Main Organization',
        `**Name:** ${user.data.organization.name}\n**SID:** ${user.data.organization.sid}\n**Rank:** ${user.data.organization.rank}`,
      );
    }

    await this.sendMessage(message.channel.id, playerEmbed);
  }
}
