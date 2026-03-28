import type { ColorValue, StyleProp } from 'react-native'
import type { TColorData, TStyle } from './types'

import { useTheme } from '@hooks'
import { isFunction, isNullish } from 'radashi'
import { StyleSheet } from 'react-native'

import { isColorName } from './isColorName'

export const useColorNames = <T>(
  // eslint-disable-next-line @typescript-eslint/prefer-readonly-parameter-types
  style: StyleProp<T>,
  colorData: TColorData<T extends TStyle ? T : never>
): T => {
  const theme = useTheme()

  if (isFunction(style)) {
    return style
  }

  if (isNullish(style)) {
    return {} as T
  }

  const flattenedStyle = StyleSheet.flatten(style)

  // @ts-expect-error Spread types may only be created from object types.
  const resultingStyle = { ...flattenedStyle } as T

  colorData.forEach(({ colorKey, defaultColorName }) => {
    // @ts-expect-error Type 'keyof (T extends TStyle ? T : never)' cannot be used to index type 'T'.
    const originalColor = resultingStyle[colorKey] as ColorValue | undefined
    const color = originalColor ?? defaultColorName

    if (isColorName(color)) {
      // @ts-expect-error Type 'keyof (T extends TStyle ? T : never)' cannot be used to index type 'T'.
      resultingStyle[colorKey] = theme.colors[color]
    }
  })

  return resultingStyle
}
