import { hide } from 'expo-splash-screen'
import { useEffect } from 'react'

import { globalInit } from './globalInit'
import { useHasHydrated } from './useHasHydrated'

globalInit()

export const useInit = (): boolean => {
  const isInitialized = [useHasHydrated()].every(Boolean)

  useEffect(() => {
    if (isInitialized) {
      hide()
    }
  }, [isInitialized])

  return isInitialized
}
