import { Message } from 'eris';
import { Command, Embed } from '../../structures/';
import { CommandObjects } from '../../@types/command';
import Redis from 'ioredis';
import { handleError } from '../../utils';

export default class info extends Command {
  readonly redis = new Redis();

  constructor(commandObjects: CommandObjects) {
    super(commandObjects, {
      name: 'info',
      description: 'Get basic bot information.',
      module: 'information',
      aliases: ['information', 'stats'],
      disabled: false,
    });
  }

  public async execute(message: Message, args: string[]) {
    await this.startTyping(message.channel.id);

    // const stats = await this.ipc.getStats();

    const statsEmbed = new Embed()
      .setTitle('Bot Stats')
      .addField(
        'Version:',
        process.env.npm_package_version,
        true,
      )
      .addField('Cluster ID:', this.clusterID, true)
      .addField('Worker ID:', this.workerID, true)
      .setTimestamp();

    this.redis
      .multi()
      .get('shardCount')
      .get('guilds')
      .get('users')
      // .get('totalRam')
      // .get('clustersRam')
      // .get('servicesRam')
      .exec()
      .then(async (responses) => {
        //  need to add proper error handling in future

        statsEmbed
          .addField('Shard Count:', responses[0][1], true)
          .addField('Server Count:', responses[1][1], true)
          .addField('User Count:', responses[2][1], true);
        // .addField(
        //   'Total Ram:',
        //   `${Math.round(responses[3][1])} MB`,
        //   true,
        // )
        // .addField(
        //   'Clusters Ram:',
        //   `${Math.round(responses[4][1])} MB`,
        //   true,
        // )
        // .addField(
        //   'Services Ram:',
        //   `${Math.round(responses[5][1])} MB`,
        //   true,
        // );

        // try and get shard
        const shard = this.bot.shards.get(
          this.workerID - 1,
        );
        // if shard isn't undefined, add api latency
        if (shard !== undefined)
          statsEmbed.addField(
            'API Latency',
            Math.round(shard.latency),
            true,
          );

        await this.sendMessage(
          message.channel.id,
          statsEmbed,
        );
      })
      .catch((err) => {
        handleError(err);
      });

    // statsEmbed
    //   .addField('Shard Count:', stats.shardCount, true)
    //   .addField('Server Count:', stats.guilds, true)
    //   .addField('User Count:', stats.users, true)
    //   .addField(
    //     'Total Ram:',
    //     `${Math.round(stats.totalRam)} MB`,
    //     true,
    //   )
    //   .addField(
    //     'Clusters Ram:',
    //     `${Math.round(stats.clustersRam)} MB`,
    //     true,
    //   )
    //   .addField(
    //     'Services Ram:',
    //     `${Math.round(stats.servicesRam)} MB`,
    //     true,
    //   )
  }
}
