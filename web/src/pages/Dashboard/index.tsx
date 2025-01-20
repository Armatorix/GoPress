import { useArticleStats } from '@/api/articles'
import { Card } from '@material-tailwind/react'

function Dashboard() {
  const stats = useArticleStats()

  if (stats.status !== 'success') {
    return (
      <div className="w-full h-20 animate-pulse bg-gray-500/50 rounded-lg" />
    )
  }
  const data = [
    {
      title: 'Total',
      value: stats.data.data.total,
    },
    {
      title: 'Published',
      value: stats.data.data.released,
    },
  ]
  return (
    <div className="w-full grid grid-cols-2 md:grid-cols-6 gap-2">
      {data.map((item) => (
        <Card
          key={item.title}
          className="bg-gray-100 hover:bg-blue-gray-200 duration-300 p-4 select-none gap-1"
        >
          <div className="text-center text-2xl font-bold">{item.title}</div>
          <div className="text-center font-bold">{item.value}</div>
        </Card>
      ))}
    </div>
  )
}

export default Dashboard
