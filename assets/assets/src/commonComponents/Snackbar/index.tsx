import { useTheme } from '@hooks'
import { snackbarModel } from '@mst'
import { observer } from 'mobx-react-lite'
import { Snackbar as RNPSnackbar } from 'react-native-paper'

import { Pressable } from '../Pressable'
import { Text } from '../Text'
import { DEFAULT_DURATION, INFINITY } from './constants'
import styles, { getSnackbarStyle, getTextStyle } from './styles'

export const Snackbar: React.FC = observer(() => {
  const theme = useTheme()
  const { hide, options } = snackbarModel

  if (!options) {
    return undefined
  }

  const { text, textType = 'info' } = options
  const isError = textType === 'error'

  const duration = isError ? INFINITY : (options.duration ?? DEFAULT_DURATION)

  return (
    <>
      <Pressable onPress={hide} style={styles.pressable} />
      <RNPSnackbar
        duration={duration}
        onDismiss={hide}
        style={getSnackbarStyle(isError, theme)}
        visible={true}
        wrapperStyle={styles.snackbarWrapper}
      >
        <Text selectable={isError} style={getTextStyle(textType)}>
          {text}
        </Text>
      </RNPSnackbar>
    </>
  )
})
