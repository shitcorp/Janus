import { Commands } from './command';
import { Message } from 'eris';

interface HandleMessageObject extends Commands {
  msg: Message;
  id: string;
}
