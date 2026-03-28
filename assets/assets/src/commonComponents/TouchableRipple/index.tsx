import { withColorNames } from '@hocs'
import { TouchableRipple as RNPTouchableRipple } from 'react-native-paper'

export const TouchableRipple = withColorNames(RNPTouchableRipple, [
  {
    colorKey: 'backgroundColor'
  },
  {
    colorKey: 'borderColor'
  }
])
