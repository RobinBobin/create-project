import { Snackbar } from '@commonComponents'
import { ThemeProvider } from '@react-navigation/native'
import { Stack } from 'expo-router'
import { PaperProvider } from 'react-native-paper'

import { useThemes } from './useThemes'

export const App: React.FC = () => {
  const { navigationTheme, paperTheme } = useThemes()

  return (
    <ThemeProvider value={navigationTheme}>
      <PaperProvider theme={paperTheme}>
        <Stack>
          <Stack.Screen name='index' />
        </Stack>
        <Snackbar />
      </PaperProvider>
    </ThemeProvider>
  )
}
