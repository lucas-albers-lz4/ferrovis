// .prettierrc.js - Ferrovis Mobile App Prettier Configuration
module.exports = {
  // Basic formatting
  semi: true,
  singleQuote: true,
  quoteProps: 'as-needed',
  trailingComma: 'es5',

  // Spacing and indentation
  tabWidth: 2,
  useTabs: false,

  // Line length
  printWidth: 80,

  // JSX specific
  jsxSingleQuote: true,
  bracketSameLine: false,

  // Other formatting
  bracketSpacing: true,
  arrowParens: 'avoid',
  endOfLine: 'lf',

  // File type overrides
  overrides: [
    {
      files: '*.json',
      options: {
        singleQuote: false,
      },
    },
  ],
};
