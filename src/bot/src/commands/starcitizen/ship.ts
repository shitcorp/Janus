import { Message } from 'eris';
import {
  Command,
  Embed,
  generalError,
} from '../../structures/';
import { CommandObjects } from '../../@types/command';
import { getShip, logger } from '../../utils';

export default class ship extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'ship',
      description: 'Get information about a specific ship.',
      module: 'starcitizen',
      aliases: [],
      args: true,
      usage: '<ship>',
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const searchTerm = args[0]; // args.join(' ');

    const ship = await getShip(searchTerm);

    logger.debug(ship);

    // If player doesn't exist or data is null
    if (
      ship === undefined ||
      ship.data === null ||
      Object.keys(ship.data).length === 0
    ) {
      await this.sendMessage(
        message.channel.id,
        generalError(
          `No ship could be found using the term \`${searchTerm}\``,
        ),
      );
      return;
    }

    const shipEmbed = new Embed()
      .setTitle(ship.data[0].name)
      .setDescription(
        ship.data[0].description.split('.')[0] + '.',
      )
      // .setThumbnail(org.data.logo)
      .setURL(
        `http://robertsspaceindustries.com${ship.data[0].url}`,
      )
      .addField('Size:', ship.data[0].size, true)
      .addField('Focus:', ship.data[0].focus, true)
      .addField('Price:', ship.data[0].price, true)
      .addField('Max crew:', ship.data[0].max_crew, true)
      .addField('Min crew:', ship.data[0].min_crew, true)
      .addField(
        'Prod Status:',
        ship.data[0].production_status,
        true,
      )
      .setTimestamp();

    // if (ship.data[0].media.length > 0) {
    //   // console.log(ship.data[0].media[0]);
    //
    //   if (
    //     ship.data[0].media[0] !== undefined &&
    //     isURL(ship.data[0].media[0].source_url)
    //   )
    //     shipEmbed.setImage(
    //       ship.data[0].media[0].source_url,
    //     );
    //   else
    //     shipEmbed.setImage(
    //       `http://robertsspaceindustries.com${ship.data[0].media[0].source_url}`,
    //     );
    // }

    await this.sendMessage(message.channel.id, shipEmbed);
  }
}
