/** @type {import("@babel/core").ConfigFunction} */
// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export default function config(api) {
  api.cache.forever()

  return {
    presets: ['babel-preset-expo']
  }
}
