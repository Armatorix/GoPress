import { DefaultService, type PostLogin } from '@/api/gen'
import { useMutation } from '@tanstack/react-query'

export function useLogin() {
  return useMutation({
    mutationFn: (body: PostLogin) => DefaultService.postAdminLogin(body),
  })
}
