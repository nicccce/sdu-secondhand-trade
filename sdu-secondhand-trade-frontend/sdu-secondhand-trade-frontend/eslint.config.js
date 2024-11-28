import js from '@eslint/js'
import pluginVue from 'eslint-plugin-vue'

export default [
  {
    name: 'app/files-to-lint',
    files: ['**/*.{js,mjs,jsx,vue}'],
  },

  {
    name: 'app/files-to-ignore',
    ignores: ['**/dist/**', '**/dist-ssr/**', '**/coverage/**'],
  },

  js.configs.recommended,
  ...pluginVue.configs['flat/essential'],

  {
    eslintConfig: {
      root: true,
      env: {
        node: true,
      },
      extends: [
        'plugin:vue/vue3-essential',
        'eslint:recommended',
      ],
      parserOptions: {
        parser: '@babel/eslint-parser',
      },
      rules: {
        'no-unused-vars': 'off',
        'vue/no-unused-components': 'off',
        'vue/multi-word-component-names': 'off',
      },
    },
  },
]
