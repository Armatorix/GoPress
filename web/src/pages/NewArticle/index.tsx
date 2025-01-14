import { useNewArticle } from '@/api/articles'
import type { PostArticle } from '@/api/gen'
import Editor from '@/components/Editor'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { Bootstrap } from '@/page_wrappers'
import { Button, Input, Switch } from '@material-tailwind/react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

function NewArticle() {
  const newArticle = useNewArticle()
  const navigate = useNavigate()
  const [form, setForm] = useState({
    title: '',
    body: '',
    description: '',
  } as PostArticle)
  const toast = useToast()
  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault()
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
        <form className="flex flex-col gap-4" onSubmit={onSubmit}>
          <Input
            type="text"
            label="Title"
            value={form.title}
            required
            onChange={(e) => setForm({ ...form, title: e.target.value })}
          />
          <div>
            <Editor
              label="Description"
              value={form.description}
              setValue={(description) => setForm({ ...form, description })}
            />
          </div>
          <div>
            <Editor
              label="Body"
              value={form.body}
              setValue={(body) => setForm({ ...form, body })}
            />
          </div>
          <Switch
            checked={form.released}
            onChange={(e) => setForm({ ...form, released: e.target.checked })}
            label="Publish"
            color="green"
          />
          <Button type="submit" loading={newArticle.isPending}>
            Submit
          </Button>
        </form>
      </div>
    </Bootstrap>
  )
}

export default NewArticle
