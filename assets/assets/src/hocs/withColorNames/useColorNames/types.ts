import type { TColor } from '@types'
import type { ImageStyle, TextStyle, ViewStyle } from 'react-native'

type TStyle = ImageStyle | TextStyle | ViewStyle
type TStyleKeys<T extends TStyle = TStyle> = T extends unknown ? keyof T : never
type TColorStyleKeys = Extract<TStyleKeys, 'color' | `${string}Color`>

interface IColorDatum<T extends TColorStyleKeys> {
  colorKey: T
  defaultColorName?: TColor
}

type TColorData<T extends TColorStyleKeys = TColorStyleKeys> =
  readonly Readonly<IColorDatum<T>>[]

export type { TColorData, TStyle }
