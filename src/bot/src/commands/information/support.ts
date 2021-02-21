import { Message } from 'eris';
import { Command, Embed } from '../../structures/';
import { CommandObjects } from '../../@types/command';

export default class support extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'support',
      description:
        'Get the link to the support & feedback server for the bot.',
      module: 'information',
      aliases: ['feedback'],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const supportEmbed = new Embed()
      .setTitle('Support & Feedback')
      .setDescription(
        "[Click here](https://top.gg/bot/776100457260384266/vote) to join Janus' support and feedback server.",
      )
      .setTimestamp();

    await this.sendMessage(
      message.channel.id,
      supportEmbed,
    );
  }
}
