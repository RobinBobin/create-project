import type { TColor } from '@types'
import type { ImageStyle, TextStyle, ViewStyle } from 'react-native'

type TStyle = ImageStyle | TextStyle | ViewStyle

interface IColorDatum<T extends TStyle> {
  colorKey: keyof T
  defaultColorName?: TColor
}

type TColorData<T extends TStyle> = readonly Readonly<IColorDatum<T>>[]

export type { IColorDatum, TColorData, TStyle }
