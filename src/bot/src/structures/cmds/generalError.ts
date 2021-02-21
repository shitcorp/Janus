import { Embed } from '../Embed';

export const generalError = (message: string): Embed => {
  return new Embed()
    .setTitle('Error')
    .setDescription(message)
    .setColor(0xcc0000)
    .setTimestamp();
};
