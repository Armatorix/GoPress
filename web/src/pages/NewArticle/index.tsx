import { useNewArticle } from '@/api/articles'
import type { PostArticle } from '@/api/gen'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { Bootstrap } from '@/page_wrappers'
import { useState } from 'react'

function NewArticle() {
  const newArticle = useNewArticle()
  const [form, setForm] = useState({
    title: '',
    body: '',
  } as PostArticle)
  const toast = useToast()
  const onSubmit = () => {
    newArticle.mutate(form, {
      onSuccess: () => {
        toast.success('Article created')
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
        <form onSubmit={onSubmit}>
          <label>
            Title
            <input name="title" />
          </label>
          <label>
            Body
            <textarea name="body" />
          </label>
          <button type="submit">Submit</button>
        </form>
      </div>
    </Bootstrap>
  )
}

export default NewArticle
