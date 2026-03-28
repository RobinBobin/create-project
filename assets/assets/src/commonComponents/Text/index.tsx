import { COLOR } from '@enums'
import { withColorNames } from '@hocs'
import { Text as RNPText } from 'react-native-paper'

export const Text = withColorNames(RNPText, [
  {
    colorKey: 'color',
    defaultColorName: COLOR.ON_SURFACE
  }
])
