import { DefaultService, type PostArticle, type PostLogin } from '@/api/gen'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'

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

export function useNewArticle() {
  const queryClient = useQueryClient()
  return useMutation({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKey.articles() })
    },
    mutationFn: (body: PostArticle) => DefaultService.createArticle(body),
  })
}

export function useUpdateArticle(id: number) {
  const queryClient = useQueryClient()
  return useMutation({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKey.articles() })
    },
    mutationFn: (body: PostArticle) => DefaultService.updateArticle(id, body),
  })
}

export function usePublishArticle() {
  const queryClient = useQueryClient()
  return useMutation({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKey.articles() })
    },
    mutationFn: (id: number) => DefaultService.publishArticle(id),
  })
}
