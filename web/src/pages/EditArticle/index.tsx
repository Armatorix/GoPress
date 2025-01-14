import { useArticles, useUpdateArticle } from '@/api/articles'
import type { Article, PatchArticle } from '@/api/gen'
import ArticleForm from '@/components/ArticleForm'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { Bootstrap } from '@/page_wrappers'
import { useNavigate, useParams } from 'react-router-dom'

function EditArticle() {
  const id = parseInt(useParams()['id'] || '')
  const { data } = useArticles()
  if (!data) return null
  const article = data.data.find((a: Article) => a.id === id)
  if (!article) return null

  return <EditWithData article={article} />
}

function EditWithData({ article }: { article: Article }) {
  const updateArticle = useUpdateArticle(article.id)
  const navigate = useNavigate()
  const toast = useToast()
  const onSubmit = (form: PatchArticle) => {
    if (!form.title || !form.body || !form.description) {
      toast.error('All fields are required')
      return
    }
    updateArticle.mutate(form, {
      onSuccess: () => {
        toast.success(`Article ${article.title} updated`)
        navigate('/articles')
      },
      onError: (error) => {
        toast.error(error.message)
      },
    })
  }
  return (
    <Bootstrap>
      <div>
        <h1>Edit Article</h1>
        <ArticleForm
          initData={article}
          onSubmit={onSubmit}
          isPending={updateArticle.isPending}
        />
      </div>
    </Bootstrap>
  )
}

export default EditArticle
