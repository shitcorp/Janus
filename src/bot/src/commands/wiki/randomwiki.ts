import { Message } from 'eris';
import {
  Command,
  Embed,
  generalError,
} from '../../structures/';
import { CommandObjects } from '../../@types/command';
import { getRandomPage } from '../../utils';
import * as Sentry from '@sentry/node';

export default class randomwiki extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'randomwiki',
      description:
        'Get a random article from the starcitizen.tools wiki.',
      module: 'wiki',
      aliases: ['randwiki'],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    const page = await getRandomPage();

    if (page === undefined || page === null) {
      await this.sendMessage(
        message.channel.id,
        generalError(
          'Oddly, no page was found. Please try again.',
        ),
      );

      Sentry.addBreadcrumb({
        category: 'cmd',
        message: 'RandomWiki util returned undefined/null.',
        level: Sentry.Severity.Warning,
      });
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
