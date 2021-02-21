import { Embed } from '../Embed';

export const properUsage = (
  commandName: string,
  usage: string | undefined,
): Embed => {
  const prefix = process.env.PREFIX;

  if (typeof usage === 'undefined') usage = '<unknown>';

  return new Embed()
    .setTitle('Proper Usage')
    .setDescription(`${prefix}${commandName} ${usage}`)
    .setColor(0xcc0000)
    .setTimestamp();
};
