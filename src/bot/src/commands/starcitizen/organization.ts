import { Message } from 'eris';
import {
  Command,
  Embed,
  generalError,
} from '../../structures/';
import { CommandObjects } from '../../@types/command';
import { getOrg } from '../../utils';

export default class organization extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'organization',
      description:
        'Get information about a specific organization.',
      module: 'starcitizen',
      aliases: ['org'],
      args: true,
      usage: '<organization SID>',
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const org = await getOrg(args[0]);

    // If player doesn't exist or data is null
    if (
      org === undefined ||
      org.data === null ||
      Object.keys(org.data).length === 0
    ) {
      await this.sendMessage(
        message.channel.id,
        generalError(
          `No org is under the SID \`${args[0]}\``,
        ),
      );
      return;
    }

    const orgEmbed = new Embed()
      .setTitle(org.data.name)
      .setThumbnail(org.data.logo)
      .setURL(org.data.url)
      .setImage(org.data.banner)
      .addField('Archetype:', org.data.archetype, true)
      .addField('Members:', org.data.members, true)
      .addField('Language:', org.data.lang, true)
      .addField(
        'Focus',
        `**Primary**: ${org.data.focus.primary.name}\n**Secondary**: ${org.data.focus.secondary.name}`,
        true,
      )
      .addField('Roleplay', org.data.roleplay, true)
      .addField('Recruiting', org.data.recruiting, true)
      .addField('Commitment:', org.data.commitment, true)
      .setTimestamp();

    await this.sendMessage(message.channel.id, orgEmbed);
  }
}
