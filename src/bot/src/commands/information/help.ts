import { Message } from 'eris';
import { Command, Embed } from '../../structures/';
import { CommandObjects } from '../../@types/command';

export default class help extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'help',
      description: 'A generic help command.',
      module: 'information',
      aliases: [],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const prefix = process.env.PREFIX;

    const helpEmbed = new Embed()
      .setTitle('Janus')
      .setDescription(
        'A fully featured and high performance bot for Star Citizen with access to a plethora of information.',
      )
      .addField(
        'Prefix:',
        `This server's prefix is \`${prefix}\``,
      )
      .addField(
        'Commands:',
        '[Click here](https://janusbot.netlify.app/#commands) to view all commands.',
      )
      .addField(
        'Invite:',
        '[Click here](https://discord.com/oauth2/authorize?client_id=776100457260384266&scope=bot&permissions=281600) to invite the bot to your server.',
      )
      .addField(
        'Support & Feedback',
        '[Click here](https://discord.gg/NVapnsA) to join the development server.',
      )
      .setTimestamp();

    await this.sendMessage(message.channel.id, helpEmbed);
  }
}
