import {
  useArticles,
  useDeleteArticle,
  usePublishArticle,
} from '@/api/articles'
import type { Article } from '@/api/gen'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { Bootstrap } from '@/page_wrappers'
import { TrashIcon } from '@heroicons/react/24/solid'
import {
  Button,
  Dialog,
  DialogBody,
  DialogFooter,
  DialogHeader,
  IconButton,
  Switch,
  Typography,
} from '@material-tailwind/react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

function Articles() {
  const articles = useArticles()
  const navigate = useNavigate()
  const onNewArticle = () => {
    navigate('/articles/new')
  }

  if (articles.isLoading || articles.data === undefined) {
    return (
      <Bootstrap>
        <Typography>Loading...</Typography>
      </Bootstrap>
    )
  }
  return (
    <Bootstrap>
      <div className="flex flex-col gap-4">
        <Button onClick={onNewArticle}>New Article</Button>
        {articles.data.data.length === 0 ? (
          <Typography>No articles found</Typography>
        ) : (
          <ArticleList articles={articles.data.data} />
        )}
      </div>
    </Bootstrap>
  )
}
type ArticleTileProps = {
  article: Article
}

function ArticleList({ articles }: { articles: Article[] }) {
  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 gap-x-2 gap-y-4 max-h-96 overflow-auto  shadow p-2">
      {articles.map((article) => (
        <ArticleTile key={article.id} article={article} />
      ))}
    </div>
  )
}
function ArticleTile({ article }: ArticleTileProps) {
  const publishArticle = usePublishArticle()
  const toast = useToast()
  const navigate = useNavigate()
  const onPublishClick = () => {
    publishArticle.mutate(article.id, {
      onSuccess: () => {
        toast.success(
          `Article ${article.title} ${
            article.released ? 'unpublished' : 'published'
          }`,
        )
      },
      onError: (error) => {
        toast.error(error.message)
      },
    })
  }
  const onEditClick = () => {
    navigate(`/articles/${article.id}/edit`)
  }
  return (
    <div className="bg-white shadow-md p-4 rounded-md h-40 relative">
      <Typography variant="h6" className="max-w-40 overflow-ellipsis">
        {article.title}
      </Typography>
      <Typography variant="small" className="overflow-hidden max-h-24">
        <div dangerouslySetInnerHTML={{ __html: article.description }} />
      </Typography>
      <div className="absolute top-4 right-4">
        <Switch
          checked={article.released}
          onChange={onPublishClick}
          label="Published"
          color="green"
        />
      </div>
      <div className="absolute bottom-4 right-4 gap-2 flex">
        <DeleteWithConfirmationButton article={article} />
        <Button onClick={onEditClick}>Edit</Button>
      </div>
    </div>
  )
}

function DeleteWithConfirmationButton({ article }: { article: Article }) {
  const deleteArticle = useDeleteArticle(article.id)
  const toast = useToast()
  const [open, setOpen] = useState(false)
  const onDelete = () => {
    deleteArticle.mutate(undefined, {
      onSuccess: () => {
        toast.success(`Article deleted`)
        setOpen(false)
      },
      onError: (error) => {
        toast.error(error.message)
      },
    })
  }
  return (
    <>
      <IconButton onClick={() => setOpen(true)} color="red">
        <TrashIcon fill="white" className="size-4" />
      </IconButton>

      <Dialog open={open} handler={() => setOpen(!open)}>
        <DialogHeader>Delete article</DialogHeader>
        <DialogBody className="flex flex-col pl-8 gap-4">
          Are you sure you want to delete this article?
          <p>{article.title}</p>
        </DialogBody>
        <DialogFooter>
          <Button
            variant="text"
            color="red"
            onClick={() => setOpen(false)}
            className="mr-1"
          >
            Cancel
          </Button>
          <Button
            variant="gradient"
            color="green"
            onClick={onDelete}
            type="submit"
          >
            Change
          </Button>
        </DialogFooter>
      </Dialog>
    </>
  )
}

export default Articles
