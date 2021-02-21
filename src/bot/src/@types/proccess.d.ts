namespace NodeJS {
  interface ProcessEnv {
    SENTRY_URL: string;
    DISCORD_TOKEN: string;
    NODE_ENV: 'development' | 'production';
    PORT: string;
    PREFIX: string;
    // set by npm when bot is started with an npm script
    npm_package_version: string;
  }
}
