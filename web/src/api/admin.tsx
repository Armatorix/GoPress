import { useMutation } from '@tanstack/react-query'
import { DefaultService, type PatchPassword } from './gen'

export function useUpdatePassword() {
  return useMutation({
    mutationFn: (body: PatchPassword) => DefaultService.updatePassword(body),
  })
}
