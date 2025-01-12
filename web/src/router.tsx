import Home from '@/pages/Home'
import Login from '@/pages/Login'
import { Route, Routes } from 'react-router-dom'

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
      <Route path="*" element={<Home />} />
    </Routes>
  )
}

export { AuthedRouter, UnauthedRouter }
