namespace NodeJS {
  interface ProcessEnv {
    SENTRY_URL: string;
    NODE_ENV: 'development' | 'production';
    // set by npm when bot is started with an npm script
    npm_package_version: string;
  }
}
