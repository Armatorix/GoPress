import Dashboard from '@/pages/Dashboard'
import Login from '@/pages/Login'
import { Navigate, Route, Routes } from 'react-router-dom'
import Articles from './pages/Articles'

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
      <Route path="articles" element={<Articles />} />
      <Route path="*" element={<Navigate to="/dashboard" />} />
    </Routes>
  )
}

export { AuthedRouter, UnauthedRouter }
