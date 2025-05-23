// babel.config.js - Babel configuration for Expo and Jest
module.exports = function (api) {
  api.cache(true);
  return {
    presets: ['babel-preset-expo'],
    plugins: [
      // Add any additional plugins if needed
    ],
  };
};
