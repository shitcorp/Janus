import { Embed } from '../Embed';

export const success = (message: string): Embed => {
  return new Embed()
    .setTitle('Success')
    .setDescription(message)
    .setColor(0x32cd32)
    .setTimestamp();
};
