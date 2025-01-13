import Dashboard from '@/pages/Dashboard'
import Login from '@/pages/Login'
import { Navigate, Route, Routes } from 'react-router-dom'

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
      <Route path="articles" element={<Navigate to="/articles" />} />
      <Route path="*" element={<Navigate to="/dashboard" />} />
    </Routes>
  )
}

export { AuthedRouter, UnauthedRouter }
