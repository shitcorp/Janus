import { EmbedOptions } from 'eris';

interface EmbedField {
  name: string;
  value: string;
  inline?: boolean;
}

// interface EmbedObject {
//   title?: string;
//   url?: string;
//   description?: string;
//   author?: EmbedAuthorOptions;
//   color?: number;
//   fields?: Embed[];
//   // image?: string;
//   thumbnail?: string;
//   timestamp?: number;
//   footer?: EmbedFooterOptions;
//   //  ^
//   readonly createdAt?: Date;
// }

/**
 * Creates an Embed
 */
export class Embed {
  private embed: EmbedOptions = {};

  /**
   * Adds a field to the embed (max 25).
   * @param name
   * @param value
   * @param inline
   */
  public addField(
    name: string,
    value: string,
    inline = false,
  ): this {
    if (!this.embed.fields) {
      this.embed.fields = [];
    }

    // if 25 or more embeds, don't add more
    if (this.embed.fields.length >= 25) return this;

    const field: EmbedField = {
      name,
      value,
      inline,
    };

    this.embed.fields.push(field);

    return this;
  }

  /**
   * Sets the author of this embed.
   * @param name
   * @param iconURl
   * @param url
   */
  public setAuthor(
    name: string,
    iconURl = '',
    url = '',
  ): this {
    this.embed.author = { name };

    if (iconURl) {
      this.embed.author.icon_url = iconURl;
    }

    if (url) {
      this.embed.author.url = url;
    }

    return this;
  }

  /**
   * Sets the color of this embed.
   * @param color
   */
  public setColor(color: number): this {
    this.embed.color = color;

    return this;
  }

  /**
   * Sets the description of this embed.
   * @param description
   */
  public setDescription(description: string): this {
    this.embed.description = description;

    return this;
  }

  /**
   * Sets the footer of this embed.
   * @param text
   * @param iconURl
   */
  public setFooter(text: string, iconURl = ''): this {
    this.embed.footer = { text };

    if (iconURl) {
      this.embed.footer.icon_url = iconURl;
    }

    return this;
  }

  /**
   * Sets the image of this embed.
   * @param url
   */
  public setImage(url: string): this {
    this.embed.url = url;

    return this;
  }

  /**
   * Sets the timestamp of this embed.
   * @param url
   */
  public setThumbnail(url: string): this {
    this.embed.thumbnail = { url };

    return this;
  }

  /**
   * Sets the title of this embed.
   * @param title
   */
  public setTitle(title: string): this {
    this.embed.title = title;

    return this;
  }

  /**
   * Sets the URL of this embed;
   * @param url
   */
  public setURl(url: string): this {
    this.embed.url = url;

    return this;
  }

  public toJSON(): EmbedOptions {
    return this.embed;
  }
}
