import { useArticles } from '@/api/articles'
import type { Article } from '@/api/gen'
import { Bootstrap } from '@/page_wrappers'
import { Typography } from '@material-tailwind/react'

function Articles() {
  const articles = useArticles()

  if (articles.isLoading || articles.data === undefined) {
    return (
      <Bootstrap>
        <Typography>Loading...</Typography>
      </Bootstrap>
    )
  }
  console.log(articles.data)
  return (
    <Bootstrap>
      {articles.data.data.length === 0 ? (
        <Typography>No articles found</Typography>
      ) : (
        <ArticleList articles={articles.data.data} />
      )}
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
      <Typography variant="small">{article.description}</Typography>
      <Typography variant="small">
        {article.released ? 'Released' : 'Not released'}
      </Typography>
    </div>
  )
}

export default Articles
