import type { SnackbarProps } from 'react-native-paper'

type TSnackbarTextType = 'error' | 'info'

type TShowOptions = Readonly<
  Pick<SnackbarProps, 'duration'> & {
    text: string
    textType?: TSnackbarTextType
  }
>

interface ISnackbarModelVolatile {
  options?: TShowOptions | undefined
}

export type { ISnackbarModelVolatile, TShowOptions, TSnackbarTextType }
