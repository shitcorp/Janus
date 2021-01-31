namespace NodeJS {
  interface ProcessEnv {
    SENTRY_URL: string;
    DISCORD_TOKEN: string;
    NODE_ENV: 'development' | 'production';
    PORT?: string;
  }
}
