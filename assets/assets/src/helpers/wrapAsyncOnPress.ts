import type { TOnPress } from '@types'

type TAsyncOnPress = () => Promise<void>

export const wrapAsyncOnPress = (asyncOnPress: TAsyncOnPress): TOnPress => {
  return () => {
    const wrapper: TAsyncOnPress = async () => {
      await asyncOnPress()
    }

    void wrapper()
  }
}
