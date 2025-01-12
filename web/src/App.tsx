import AuthProvider from '@/lib/AuthProvider'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { jwtDecode } from 'jwt-decode'
import { useEffect, useState } from 'react'
import { BrowserRouter } from 'react-router-dom'
import { ApiError, OpenAPI } from './api/gen'
import { AuthedRouter, UnauthedRouter } from './router'
import { BACKEND_URL } from './lib/config'
import { ToastProvider } from './components/Toasts/ToastsProvider'

OpenAPI.BASE = `${BACKEND_URL}/api/v1`
function App() {
  const [isAuthed, setIsAuthed] = useState(false)
  useEffect(() => {
    try {
      const jwt = localStorage.getItem('jwt') || ''
      const jwtDecoded = jwtDecode(jwt)
      const expAt = jwtDecoded.exp || 0
      const now = new Date().getTime() / 1000
      if (expAt > now) {
        setIsAuthed(true)
        OpenAPI.WITH_CREDENTIALS = true
        OpenAPI.TOKEN = jwt
      } else {
        localStorage.removeItem('jwt')
        setIsAuthed(false)
        OpenAPI.WITH_CREDENTIALS = false
        OpenAPI.TOKEN = undefined
      }
    } catch (e) {
      setIsAuthed(false)
      OpenAPI.WITH_CREDENTIALS = false
      OpenAPI.TOKEN = undefined
    }
  }, [isAuthed])

  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 1 * 60 * 1000,
        retry: (failureCount: number, error: unknown) => {
          const maxRetries = 2
          if ((error as ApiError)?.status === 401) {
            localStorage.removeItem('jwt')
            setIsAuthed(false)
            return false
          }
          return (error as ApiError)?.status >= 500 && failureCount < maxRetries
        },
      },
    },
  })

  return (
    <BrowserRouter>
      <AuthProvider>
        <QueryClientProvider client={queryClient}>
          <ToastProvider>
            {isAuthed ? <AuthedRouter /> : <UnauthedRouter />}
          </ToastProvider>
        </QueryClientProvider>
      </AuthProvider>
    </BrowserRouter>
  )
}

export default App
