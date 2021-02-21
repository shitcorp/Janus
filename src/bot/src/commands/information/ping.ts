import { Message } from 'eris';
import { Command, Embed, success } from '../../structures/';
import { CommandObjects } from '../../@types/command';
import { formatMessage, handleError } from '../../utils';

export default class ping extends Command {
  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'ping',
      description: 'Sends basic latency info.',
      module: 'information',
      aliases: [],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    const pingEmbed = new Embed().setTitle('Ping?');

    const m = await this.sendMessage(
      message.channel.id,
      pingEmbed,
    );

    if (m instanceof Message) {
      // try and get shard
      const shard = this.bot.shards.get(this.workerID - 1);

      // if shard is undefined, return
      if (shard === undefined) return;

      await m
        .edit(
          formatMessage(
            success(
              `Pong! Latency is ${
                m.createdAt - message.createdAt
              }ms. API Latency is ${Math.round(
                shard.latency,
              )}ms`,
            ),
          ),
        )
        .catch((err) => {
          handleError(err);
        });
    }
  }
}
