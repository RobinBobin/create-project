import { withColorNames } from '@hocs'
import { Pressable as RNPressable } from 'react-native'

export const Pressable = withColorNames(RNPressable, [
  {
    colorKey: 'backgroundColor'
  }
])
