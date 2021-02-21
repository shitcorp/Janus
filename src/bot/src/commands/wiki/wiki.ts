import { Message } from 'eris';
import {
  Command,
  Embed,
  generalError,
} from '../../structures/';
import { CommandObjects } from '../../@types/command';
import { getPage } from '../../utils';

export default class wiki extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'wiki',
      description:
        'Get an article from the starcitizen.tools wiki.',
      module: 'wiki',
      aliases: [],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const searchTerm = args.join(' ');
    const page = await getPage(searchTerm);

    if (page === undefined || page === null) {
      await this.sendMessage(
        message.channel.id,
        generalError('No pages were found under that term'),
      );
      return;
    }

    const wikiEmbed = new Embed()
      .setTitle(page.title)
      .setURL(page.fullURl)
      .setAuthor(
        'Star Citizen Wiki',
        'https://starcitizen.tools/resources/assets/apple-touch-icon.png',
      )
      .setDescription(page.summary)
      .setThumbnail(page.mainImage)
      .setFooter(
        'Content by starcitizen.tools',
        'https://starcitizen.tools/resources/assets/apple-touch-icon.png',
      )
      .setTimestamp();

    await this.sendMessage(message.channel.id, wikiEmbed);
  }
}
