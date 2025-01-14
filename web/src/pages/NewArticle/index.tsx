import { useNewArticle } from '@/api/articles'
import type { PostArticle } from '@/api/gen'
import Editor from '@/components/Editor'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { useAuthInfo } from '@/lib/AuthProvider'
import { Bootstrap } from '@/page_wrappers'
import { Button } from '@material-tailwind/react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

function NewArticle() {
  const newArticle = useNewArticle()
  const navigate = useNavigate()
  const [form, setForm] = useState({
    title: '',
    body: '',
  } as PostArticle)
  const toast = useToast()
  const onSubmit = () => {
    newArticle.mutate(form, {
      onSuccess: () => {
        navigate('/articles')
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
        <form>
          <div>
            <Editor
              value={form.body}
              setValue={(body) => setForm({ ...form, body })}
            />
          </div>

          <Button type="submit" onClick={onSubmit}>
            Submit
          </Button>
        </form>
      </div>
    </Bootstrap>
  )
}

export default NewArticle
