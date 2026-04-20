import type { StyleProp } from 'react-native'
import type { TColorData, TStyle } from './useColorNames'

import { forwardRef } from 'react'

import { useColorNames } from './useColorNames'

export const withColorNames = <
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  TWrappedComponent extends React.JSXElementConstructor<any>
>(
  WrappedComponent: TWrappedComponent,
  colorData: TColorData
): ReturnType<
  typeof forwardRef<
    React.ComponentRef<TWrappedComponent>,
    React.ComponentPropsWithRef<TWrappedComponent>
  >
> => {
  const ComponentWithRef = forwardRef<
    React.ComponentRef<TWrappedComponent>,
    React.ComponentPropsWithRef<TWrappedComponent>
  >(({ style, ...props }, ref): React.ReactElement => {
    const flattenedStyle = useColorNames(style as StyleProp<TStyle>, colorData)

    // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment, @typescript-eslint/no-explicit-any
    const Component = WrappedComponent as any

    return <Component {...props} ref={ref} style={flattenedStyle} />
  })

  return ComponentWithRef
}
