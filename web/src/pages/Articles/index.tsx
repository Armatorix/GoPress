import { useArticles, usePublishArticle } from '@/api/articles'
import type { Article } from '@/api/gen'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { Bootstrap } from '@/page_wrappers'
import { Button, Switch, Typography } from '@material-tailwind/react'
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
    <div className="grid grid-cols-2 gap-2">
      {articles.map((article) => (
        <ArticleTile key={article.id} article={article} />
      ))}
    </div>
  )
}
function ArticleTile({ article }: ArticleTileProps) {
  const publishArticle = usePublishArticle()
  const toast = useToast()
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
    </div>
  )
}

export default Articles
