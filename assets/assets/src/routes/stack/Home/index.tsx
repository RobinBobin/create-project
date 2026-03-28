import { View, Screen, Text } from '@commonComponents'
import { COLOR } from '@enums'
import React from 'react'

export const Home: React.FC = () => {
  return (
    <Screen>
      <View
        style={{
          alignItems: 'center',
          backgroundColor: COLOR.ERROR,
          flex: 1,
          justifyContent: 'center'
        }}
      >
        <Text>Edit me.</Text>
      </View>
    </Screen>
  )
}
