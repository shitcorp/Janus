import { Embed } from '../structures';
import { AdvancedMessageContent } from 'eris';

/**
 * Takes embed and converts it to the AdvancedMessageContent format
 * @param embed
 */
export const formatMessage = (
  embed: Embed,
): AdvancedMessageContent => {
  return {
    embed: {
      ...embed.toJSON(),
    },
  };
};
