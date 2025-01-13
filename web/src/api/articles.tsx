import { DefaultService, type PostLogin } from '@/api/gen'
import { useMutation, useQuery } from '@tanstack/react-query'

export const queryKey = {
  articles: () => ['articles'] as const,
}
export function useLogin() {
  return useMutation({
    mutationFn: (body: PostLogin) => DefaultService.postAdminLogin(body),
  })
}

export function useArticles() {
  return useQuery({
    queryKey: queryKey.articles(),
    queryFn: () => DefaultService.getArticles(),
  })
}
