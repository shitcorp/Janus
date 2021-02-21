import { Embed } from '../structures';
import {
  AdvancedMessageContent,
  Client,
  Message,
} from 'eris';
import { handleError } from './handleError';
import { formatMessage } from './formatMessage';

/**
 * Handles message sending
 * @param bot
 * @param channelID
 * @param message
 */
export const sendMessage = async (
  bot: Client,
  channelID: string,
  message: string | Embed | AdvancedMessageContent,
): Promise<Message | void> => {
  // if embed
  if (message instanceof Embed) {
    // return embed as json for Eris to handle
    return await bot
      .createMessage(channelID, formatMessage(message))
      .catch((err) => {
        handleError(err);
      });
  } else {
    return await bot
      .createMessage(channelID, message)
      .catch((err) => {
        handleError(err);
      });
  }
};
