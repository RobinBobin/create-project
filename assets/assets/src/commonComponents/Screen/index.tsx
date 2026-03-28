import type { PropsWithChildren } from 'react'

import { View } from '../View'
import styles from './styles'

export const Screen: React.FC<PropsWithChildren> = ({ children }) => {
  return <View style={styles.container}>{children}</View>
}
