import { isError, isString } from 'radashi'

export const getErrorMessage = (error: unknown): string => {
  if (isString(error)) {
    return error
  }

  if (isError(error)) {
    return error.message
  }

  return `An unknown error of type '${typeof error}':\n${String(error)}`
}
