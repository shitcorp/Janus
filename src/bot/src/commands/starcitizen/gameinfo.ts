import { Message } from 'eris';
import { Command, Embed } from '../../structures/';
import { CommandObjects } from '../../@types/command';
import Redis from 'ioredis';
import { handleError } from '../../utils';

export default class gameinfo extends Command {
  readonly redis = new Redis();

  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'gameinfo',
      description:
        'Get general information about StarCitizen.',
      module: 'starcitizen',
      aliases: ['scinfo', 'scstats', 'gamestats'],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const statsEmbed = new Embed().setTitle('Game Info');

    this.redis
      .multi()
      .get('current_live')
      .get('current_ptu')
      .get('fans')
      .get('fleet')
      .get('funds')
      .exec()
      .then(async (responses) => {
        //  need to add proper error handling in future

        // logger.debug(responses[0][1]);
        // logger.debug(responses[1][1]);
        // logger.debug(responses[2][1]);
        // logger.debug(responses[3][1]);
        // logger.debug(responses[4][1]);

        statsEmbed
          .addField(
            'Live:',
            responses[0][1] || 'undefined',
            true,
          )
          .addField(
            'PTU:',
            responses[1][1] || 'undefined',
            true,
          )
          .addField(
            'Fans:',
            responses[2][1] || 'undefined',
            true,
          )
          .addField(
            'Fleet:',
            responses[3][1] || 'undefined',
            true,
          )
          .addField(
            'Funds:',
            responses[4][1] || 'undefined',
            true,
          )
          .setTimestamp();

        await this.sendMessage(
          message.channel.id,
          statsEmbed,
        );
      })
      .catch((err) => {
        handleError(err);
      });
  }
}
