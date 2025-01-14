import { useArticles } from '@/api/articles'
import type { Article } from '@/api/gen'
import { Bootstrap } from '@/page_wrappers'
import { Button, Typography } from '@material-tailwind/react'
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
  return (
    <div className="bg-white shadow-md p-4 rounded-md">
      <Typography variant="h6">{article.title}</Typography>
      <Typography variant="small">
        <div dangerouslySetInnerHTML={{ __html: article.description }} />
      </Typography>
      <Typography variant="small">
        {article.released ? 'Released' : 'Not released'}
      </Typography>
    </div>
  )
}

export default Articles
