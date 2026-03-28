import { withColorNames } from '@hocs'
import { View as RNView } from 'react-native'

export const View = withColorNames(RNView, [
  {
    colorKey: 'backgroundColor'
  },
  {
    colorKey: 'borderColor'
  }
])
