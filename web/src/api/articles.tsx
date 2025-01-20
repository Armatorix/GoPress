import {
  DefaultService,
  type PostArticle,
  type PostArticleGenerate,
  type PostLogin,
} from '@/api/gen'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'

export const queryKey = {
  articles: () => ['articles'] as const,
  articlesStats: () => ['articles', 'stats'] as const,
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
      queryClient.invalidateQueries({ queryKey: queryKey.articlesStats() })
    },
    mutationFn: (body: PostArticle) => DefaultService.createArticle(body),
  })
}

export function useUpdateArticle(id: number) {
  const queryClient = useQueryClient()
  return useMutation({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKey.articles() })
      queryClient.invalidateQueries({ queryKey: queryKey.articlesStats() })
    },
    mutationFn: (body: PostArticle) => DefaultService.updateArticle(id, body),
  })
}

export function usePublishArticle() {
  const queryClient = useQueryClient()
  return useMutation({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKey.articles() })
      queryClient.invalidateQueries({ queryKey: queryKey.articlesStats() })
    },
    mutationFn: (id: number) => DefaultService.publishArticle(id),
  })
}

export function useDeleteArticle(id: number) {
  const queryClient = useQueryClient()
  return useMutation({
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: queryKey.articles() })
      queryClient.invalidateQueries({ queryKey: queryKey.articlesStats() })
    },
    mutationFn: () => DefaultService.deleteArticles([id]),
  })
}

export function useAiGenerate() {
  return useMutation({
    mutationFn: (body: PostArticleGenerate) =>
      DefaultService.generateArticle(body),
  })
}

export function useArticleStats() {
  return useQuery({
    queryKey: queryKey.articlesStats(),
    queryFn: () => DefaultService.getArticleStats(),
  })
}
