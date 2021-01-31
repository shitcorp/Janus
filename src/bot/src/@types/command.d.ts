import { Command } from '../structures';

interface CommandOptions {
  readonly name: string;
  readonly description: string;
  readonly module: string;
  readonly permissions?: string[];
  readonly aliases?: Array<string>;
  readonly args?: boolean;
  readonly usage?: string;
  readonly disabled?: boolean;
  readonly cooldown?: 3 | 5 | 7 | 10;
  readonly permlevel?: 0 | 1 | 2 | 3;
}

interface CommandImport {
  default: Command;
}
