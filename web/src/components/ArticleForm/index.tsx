import { useAiGenerate } from '@/api/articles'
import {
  ApiError,
  HtmlStyleType,
  type Article,
  type PatchArticle,
  type PostArticle,
  type PostArticleGenerate,
} from '@/api/gen'
import Editor from '@/components/Editor'
import { useToast } from '@/components/Toasts/ToastsProvider'
import {
  Button,
  Dialog,
  DialogBody,
  DialogFooter,
  DialogHeader,
  Input,
  Switch,
} from '@material-tailwind/react'
import { useState } from 'react'

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
      <AIGenerateContent setForm={setForm} form={form} />
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

type AIGenerateContentProps = {
  form: FormType
  setForm: React.Dispatch<React.SetStateAction<FormType>>
}
function AIGenerateContent({ form, setForm }: AIGenerateContentProps) {
  const [open, setOpen] = useState(false)
  const generate = useAiGenerate()
  const toast = useToast()
  const handleOpen = () => setOpen(!open)
  const [prompt, setPrompt] = useState({
    topic: '',
    htmlStyles: HtmlStyleType.TAILWIND_CSS,
  } as PostArticleGenerate)
  const handleConfirm = () => {
    generate.mutate(prompt, {
      onSuccess: (v) => {
        setForm({
          ...form,
          title: v.data.title,
          body: v.data.body,
          description: v.data.description,
        } as FormType)
        toast.success('Article generated')
        setOpen(false)
      },
      onError: (error) => {
        if (error instanceof ApiError) {
          toast.error(error.message)
          return
        } else {
          toast.error('Something went wrong')
        }
      },
    })
  }
  return (
    <>
      <Button onClick={handleOpen} variant="gradient">
        AI Generate
      </Button>
      <Dialog open={open} handler={handleOpen}>
        <DialogHeader>Generate Content with AI</DialogHeader>
        <DialogBody className="flex flex-col pl-8 gap-4">
          <Input
            type="text"
            label="Article topic"
            variant="outlined"
            value={prompt.topic}
            required
            onChange={(e) => setPrompt({ ...prompt, topic: e.target.value })}
          />
          <Switch
            checked={prompt.htmlStyles === HtmlStyleType.TAILWIND_CSS}
            onChange={(e) =>
              setPrompt({
                ...prompt,
                htmlStyles: e.target.checked
                  ? HtmlStyleType.TAILWIND_CSS
                  : HtmlStyleType.RAW_HTML,
              })
            }
            label={
              prompt.htmlStyles === HtmlStyleType.TAILWIND_CSS
                ? 'Tailwind CSS'
                : 'Raw HTML'
            }
            color="green"
          />
        </DialogBody>
        <DialogFooter>
          <Button
            variant="text"
            color="red"
            onClick={handleOpen}
            className="mr-1"
          >
            Cancel
          </Button>
          <Button
            variant="gradient"
            color="green"
            onClick={handleConfirm}
            loading={generate.isPending}
          >
            Generate
          </Button>
        </DialogFooter>
      </Dialog>
    </>
  )
}

export default ArticleForm
