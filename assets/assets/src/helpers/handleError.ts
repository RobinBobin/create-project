import { snackbarModel } from '@mst'

import { getErrorMessage } from './getErrorMessage'

export const handleError = (error: unknown): void => {
  const text = getErrorMessage(error)

  snackbarModel.show({ text, textType: 'error' })
}
