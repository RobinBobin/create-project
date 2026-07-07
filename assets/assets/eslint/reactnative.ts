import { defineConfig } from 'eslint/config'

import { ts } from './ruleOptions/index.js'

export default defineConfig([
  {
    rules: {
      '@typescript-eslint/no-shadow': [
        'error',
        {
          ...ts.noShadow,
          allow: ['Screen', 'StyleSheet', 'Text']
        }
      ]
    }
  }
])
