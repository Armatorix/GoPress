import Dashboard from '@/pages/Dashboard'
import Login from '@/pages/Login'
import { Navigate, Route, Routes } from 'react-router-dom'
import Articles from './pages/Articles'
import NewArticle from './pages/NewArticle'

function UnauthedRouter() {
  return (
    <Routes>
      <Route path="*" element={<Login />} />
    </Routes>
  )
}

function AuthedRouter() {
  return (
    <Routes>
      <Route path="dashboard" element={<Dashboard />} />
      <Route path="articles">
        <Route path="new" element={<NewArticle />} />
        <Route index element={<Articles />} />
      </Route>
      <Route path="*" element={<Navigate to="/dashboard" />} />
    </Routes>
  )
}

export { AuthedRouter, UnauthedRouter }
