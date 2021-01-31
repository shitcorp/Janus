import { Client, Message } from 'eris';
import { Command, Embed } from '../../structures/';

export default class Info extends Command {
  constructor(bot: Client) {
    super(bot, {
      name: 'info',
      description: 'test',
      module: 'info',
      aliases: ['example2'],
      disabled: true,
    });
  }

  public async execute(message: Message, args: string[]) {
    const embed = new Embed().setTitle('hello');

    await this.sendMessage(message.channel, embed);
  }
}
