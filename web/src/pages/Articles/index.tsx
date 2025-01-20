import {
  useArticles,
  useDeleteArticle,
  usePublishArticle,
} from '@/api/articles'
import type { Article } from '@/api/gen'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { TrashIcon } from '@heroicons/react/24/solid'
import {
  Button,
  Chip,
  Dialog,
  DialogBody,
  DialogFooter,
  DialogHeader,
  IconButton,
  Switch,
  Tooltip,
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
    return <Typography>Loading...</Typography>
  }
  return (
    <div className="flex flex-col gap-4">
      <Button onClick={onNewArticle}>New Article</Button>
      {articles.data.data.length === 0 ? (
        <Typography>No articles found</Typography>
      ) : (
        <ArticleList articles={articles.data.data} />
      )}
    </div>
  )
}
type ArticleTileProps = {
  article: Article
}

function ArticleList({ articles }: { articles: Article[] }) {
  return (
    <div className="grid grid-cols-1 lg:grid-cols-2 gap-x-2 gap-y-4 max-h-96 overflow-auto pb-2">
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
    <div className="bg-white shadow-md p-4 rounded-md h-52 relative hover:bg-gray-500 hover:text-white group">
      <Tooltip placement="top" content={article.title}>
        <Typography
          variant="h6"
          className="max-w-72 text-nowrap overflow-hidden text-ellipsis"
        >
          {article.title}
        </Typography>
      </Tooltip>
      <div
        className="text-sm overflow-hidden max-h-20"
        dangerouslySetInnerHTML={{ __html: article.description }}
      />
      <div className="absolute top-4 right-4">
        <Switch
          checked={article.released}
          onChange={onPublishClick}
          label="Published"
          color="green"
          labelProps={{ className: 'group-hover:text-white' }}
        />
      </div>
      <div className="absolute bottom-2 right-2 ">
        <div className="flex flex-row gap-1 items-center text-right bg-gray-500/20 group-hover:bg-white/70 group-hover:text-black p-2 rounded-md">
          <div className="flex flex-col gap-1">
            <Typography variant="small" className="font-medium text-center">
              Created
            </Typography>
            <Chip value={new Date(article.createdAt).toLocaleDateString()} />
          </div>
          <div className="flex flex-col gap-1">
            <Typography variant="small" className="font-medium text-center">
              Updated
            </Typography>
            <Chip value={new Date(article.updatedAt).toLocaleDateString()} />
          </div>
          <DeleteWithConfirmationButton article={article} />
          <Button onClick={onEditClick}>Edit</Button>
        </div>
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
