import type { TSnackbarTextType } from '@mst'
import type { TAppTheme } from '@types'
import type { StyleProp, TextStyle, ViewStyle } from 'react-native'

import { COLOR } from '@enums'
import { StyleSheet } from 'react-native'

const getSnackbarStyle = (
  isError: boolean,
  theme: TAppTheme
): StyleProp<ViewStyle> => [
  isError && {
    backgroundColor: theme.colors.errorContainer
  }
]

const getTextStyle = (textType: TSnackbarTextType): StyleProp<TextStyle> => [
  {
    fontSize: 20
  },
  textType === 'error' && { color: COLOR.ON_ERROR_CONTAINER },
  textType === 'info' && { color: COLOR.INVERSE_ON_SURFACE }
]

export { getSnackbarStyle, getTextStyle }

export default StyleSheet.create({
  pressable: {
    ...StyleSheet.absoluteFillObject,
    cursor: 'auto'
  },
  snackbarWrapper: {
    alignItems: 'center',
    justifyContent: 'center',
    top: 0
  }
})
