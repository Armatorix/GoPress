import { useState } from 'react'
import { Button, Input, Switch } from '@material-tailwind/react'
import Editor from '@/components/Editor'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { type Article, type PatchArticle, type PostArticle } from '@/api/gen'

type FormType = PostArticle | PatchArticle
interface ArticleFormProps {
  initData: Article
  onSubmit: (form: FormType) => void
  isPending: boolean
}

const ArticleForm: React.FC<ArticleFormProps> = ({
  initData,
  onSubmit,
  isPending,
}) => {
  const [form, setForm] = useState(initData as FormType)
  const toast = useToast()

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (!form.title || !form.body || !form.description) {
      toast.error('All fields are required')
      return
    }
    onSubmit(form)
  }

  return (
    <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
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
      <Button type="submit" loading={isPending}>
        Submit
      </Button>
    </form>
  )
}

export default ArticleForm
