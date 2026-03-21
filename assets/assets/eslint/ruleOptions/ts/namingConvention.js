/* eslint-disable sort-keys */

export const namingConvention = [
  {
    selector: 'interface',
    format: ['PascalCase'],
    prefix: ['I']
  },
  {
    selector: 'typeAlias',
    format: ['PascalCase'],
    prefix: ['T']
  },
  {
    selector: 'variable',
    types: ['boolean'],
    format: ['PascalCase'],
    prefix: [
      'are',
      'ARE_',
      'is',
      'IS_',
      'should',
      'SHOULD_',
      'has',
      'HAS_',
      'have',
      'HAVE_',
      'can',
      'CAN_',
      'did',
      'DID_',
      'will',
      'WILL_'
    ]
  }
]
