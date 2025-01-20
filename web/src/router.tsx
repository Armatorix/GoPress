import Dashboard from '@/pages/Dashboard'
import Login from '@/pages/Login'
import { Navigate, Outlet, Route, Routes } from 'react-router-dom'
import Articles from './pages/Articles'
import NewArticle from './pages/NewArticle'
import EditArticle from './pages/EditArticle'
import Profile from './pages/Profile'
import { Bootstrap } from './page_wrappers'

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
      <Route
        element={
          <Bootstrap>
            <Outlet />
          </Bootstrap>
        }
      >
        <Route path="dashboard" element={<Dashboard />} />
        <Route path="articles">
          <Route path="new" element={<NewArticle />} />
          <Route index element={<Articles />} />
          <Route path=":id">
            <Route path="edit" element={<EditArticle />} />
          </Route>
        </Route>
        <Route path="profile" element={<Profile />} />
        <Route path="*" element={<Navigate to="/articles" />} />
      </Route>
    </Routes>
  )
}

export { AuthedRouter, UnauthedRouter }
