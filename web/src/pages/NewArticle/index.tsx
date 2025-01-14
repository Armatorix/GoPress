import { useNewArticle } from '@/api/articles'
import type { Article, PostArticle } from '@/api/gen'
import ArticleForm from '@/components/ArticleForm'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { Bootstrap } from '@/page_wrappers'
import { useNavigate } from 'react-router-dom'

function NewArticle() {
  const newArticle = useNewArticle()
  const navigate = useNavigate()
  const toast = useToast()
  const onSubmit = (form: PostArticle) => {
    if (!form.title || !form.body || !form.description) {
      toast.error('All fields are required')
      return
    }
    newArticle.mutate(form, {
      onSuccess: () => {
        toast.success('Article created')
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
        <h1>New Article</h1>
        <ArticleForm
          initData={{} as Article}
          onSubmit={onSubmit}
          isPending={newArticle.isPending}
        />
      </div>
    </Bootstrap>
  )
}

export default NewArticle
