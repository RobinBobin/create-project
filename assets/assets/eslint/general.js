import eslintJs from '@eslint/js'
import stylisticPlugin from '@stylistic/eslint-plugin'
import tsParser from '@typescript-eslint/parser'
import { defineConfig, globalIgnores } from 'eslint/config'
import eslintPluginImportX from 'eslint-plugin-import-x'
// @ts-expect-error Could not find a declaration file for module 'eslint-plugin-promise'.
import eslintPluginPromise from 'eslint-plugin-promise'
import eslintPluginSimpleImportSort from 'eslint-plugin-simple-import-sort'
import { configs as typescriptEslintConfigs } from 'typescript-eslint'

import {
  importX,
  js,
  simpleImportSort,
  stylistic,
  ts
} from './ruleOptions/index.js'

export default defineConfig([
  eslintJs.configs.recommended,
  // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access
  eslintPluginPromise.configs['flat/recommended'],
  globalIgnores(['js/']),
  typescriptEslintConfigs.strictTypeChecked,
  typescriptEslintConfigs.stylisticTypeChecked,
  {
    extends: ['import-x/flat/recommended', 'import-x/flat/typescript'],
    files: ['**/*.{js,jsx,mjs,cjs,ts,tsx,mts,cts}'],
    languageOptions: {
      ecmaVersion: 'latest',
      parser: tsParser,
      parserOptions: {
        projectService: true
      },
      sourceType: 'module'
    },
    plugins: {
      '@stylistic': stylisticPlugin,
      'import-x': eslintPluginImportX,
      'simple-import-sort': eslintPluginSimpleImportSort
    },
    rules: {
      // @eslint/js
      'accessor-pairs': 'error',
      'array-callback-return': ['error', js.arrayCallbackReturn],
      camelcase: 'error',
      complexity: 'error',
      'consistent-return': 'error',
      'consistent-this': 'error',
      curly: 'error',
      'default-case-last': 'error',
      'dot-notation': 'error',
      eqeqeq: 'error',
      'func-name-matching': 'error',
      'func-names': 'error',
      'grouped-accessor-pairs': ['error', 'getBeforeSet'],
      'guard-for-in': 'error',
      'id-length': ['error', js.idLength],
      'max-depth': 'error',
      'max-lines': 'error',
      'max-lines-per-function': ['error', js.maxLinesPerFunction],
      'max-nested-callbacks': 'error',
      'max-statements': ['error', js.maxStatements],
      'new-cap': 'error',
      'no-array-constructor': 'error',
      'no-await-in-loop': 'error',
      'no-bitwise': 'error',
      'no-constructor-return': 'error',
      'no-div-regex': 'error',
      'no-else-return': 'error',
      'no-eval': 'error',
      'no-extend-native': 'error',
      'no-extra-bind': 'error',
      'no-extra-boolean-cast': 'error',
      'no-extra-label': 'error',
      'no-implicit-coercion': 'error',
      'no-implicit-globals': 'error',
      'no-implied-eval': 'error',
      'no-inline-comments': 'error',
      'no-invalid-this': 'error',
      'no-label-var': 'error',
      'no-labels': 'error',
      'no-lone-blocks': 'error',
      'no-lonely-if': 'error',
      'no-multi-assign': 'error',
      'no-multi-str': 'error',
      'no-nested-ternary': 'error',
      'no-new': 'error',
      'no-new-func': 'error',
      'no-new-wrappers': 'error',
      'no-object-constructor': 'error',
      'no-param-reassign': 'error',
      'no-promise-executor-return': 'error',
      'no-return-assign': 'error',
      'no-self-compare': 'error',
      'no-sequences': 'error',
      'no-template-curly-in-string': 'error',
      'no-throw-literal': 'error',
      'no-unmodified-loop-condition': 'error',
      'no-unneeded-ternary': 'error',
      'no-unreachable-loop': 'error',
      'no-unused-expressions': 'error',
      'no-useless-assignment': 'error',
      'no-useless-call': 'error',
      'no-useless-computed-key': 'error',
      'no-useless-concat': 'error',
      'no-useless-rename': 'error',
      'no-useless-return': 'error',
      'no-var': 'error',
      'object-shorthand': 'error',
      'prefer-arrow-callback': 'error',
      'prefer-object-has-own': 'error',
      'prefer-promise-reject-errors': 'error',
      'prefer-template': 'error',
      'require-atomic-updates': 'error',
      'require-await': 'error',
      'sort-keys': ['error', 'asc', js.sortKeys],
      'symbol-description': 'error',
      yoda: 'error',

      // @stylistic
      '@stylistic/jsx-curly-brace-presence': [
        'error',
        stylistic.jsxCurlyBracePresence
      ],
      '@stylistic/jsx-pascal-case': 'error',
      '@stylistic/jsx-sort-props': ['error', stylistic.jsxSortProps],

      // @typescript-eslint
      '@typescript-eslint/class-methods-use-this': 'error',
      '@typescript-eslint/consistent-type-imports': 'error',
      '@typescript-eslint/default-param-last': 'error',
      '@typescript-eslint/explicit-function-return-type': 'error',
      '@typescript-eslint/max-params': ['error', ts.maxParams],
      '@typescript-eslint/method-signature-style': 'error',
      '@typescript-eslint/naming-convention': ['error', ...ts.namingConvention],
      '@typescript-eslint/no-confusing-void-expression': [
        'error',
        ts.noConfusingVoidExpression
      ],
      '@typescript-eslint/no-loop-func': 'error',
      '@typescript-eslint/no-magic-numbers': ['error', ts.noMagicNumbers],
      '@typescript-eslint/no-shadow': ['error', ts.noShadow],
      '@typescript-eslint/no-unnecessary-parameter-property-assignment':
        'error',
      '@typescript-eslint/no-unnecessary-qualifier': 'error',
      '@typescript-eslint/no-unused-vars': ['error', ts.noUnusedVars],
      '@typescript-eslint/no-use-before-define': [
        'error',
        ts.noUseBeforeDefine
      ],
      '@typescript-eslint/no-useless-empty-export': 'error',
      '@typescript-eslint/prefer-readonly': 'error',
      '@typescript-eslint/prefer-readonly-parameter-types': 'error',
      '@typescript-eslint/restrict-template-expressions': [
        'error',
        ts.restrictTemplateExpressions
      ],
      '@typescript-eslint/strict-boolean-expressions': 'error',
      '@typescript-eslint/switch-exhaustiveness-check': 'error',

      // import-x
      'import-x/consistent-type-specifier-style': 'error',
      'import-x/exports-last': 'error',
      'import-x/first': 'error',
      'import-x/group-exports': 'error',
      'import-x/newline-after-import': 'error',
      'import-x/no-cycle': 'error',
      'import-x/no-deprecated': 'error',
      'import-x/no-duplicates': 'error',
      'import-x/no-empty-named-blocks': 'error',
      'import-x/no-extraneous-dependencies': 'error',
      'import-x/no-mutable-exports': 'error',
      'import-x/no-self-import': 'error',
      'import-x/no-unassigned-import': 'error',
      'import-x/no-unused-modules': ['error', importX.noUnusedModules],
      'import-x/no-useless-path-segments': 'error',

      //promise
      'promise/no-multiple-resolved': 'error',
      'promise/prefer-await-to-callbacks': 'error',
      'promise/prefer-await-to-then': 'error',
      'promise/spec-only': 'error',

      // simple-import-sort
      'simple-import-sort/exports': 'error',
      'simple-import-sort/imports': ['error', simpleImportSort.imports]
    }
  }
])
