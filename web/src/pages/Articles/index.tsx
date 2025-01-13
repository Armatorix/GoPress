import { useArticles } from '@/api/articles'
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
  return <Bootstrap>xd</Bootstrap>
}

export default Articles
