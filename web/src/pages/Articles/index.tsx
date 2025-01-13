import { useArticles } from '@/api/articles'
import { Bootstrap } from '@/page_wrappers'
import { Typography } from '@material-tailwind/react'

function Articles() {
  const articles = useArticles()

  if (articles.isLoading) {
    return (
      <Bootstrap>
        <Typography>Loading...</Typography>
      </Bootstrap>
    )
  }
  return (
    <Bootstrap>
      <Typography color="gray">Hello, world!</Typography>
    </Bootstrap>
  )
}

export default Articles
