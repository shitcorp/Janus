import { Message } from 'eris';
import { Command, Embed } from '../../structures/';
import { CommandObjects } from '../../@types/command';

export default class vote extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'vote',
      description: 'Get the link to vote for the bot.',
      module: 'information',
      aliases: [],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const voteEmbed = new Embed()
      .setTitle('Vote')
      .setDescription(
        '[Click here](https://top.gg/bot/776100457260384266/vote) to vote for Janus.',
      )
      .setTimestamp();

    await this.sendMessage(message.channel.id, voteEmbed);
  }
}
