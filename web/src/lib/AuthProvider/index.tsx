import { jwtDecode } from 'jwt-decode'
import React, { useEffect } from 'react'
import { useLocation, useNavigate } from 'react-router'
function useQuery() {
  const { search } = useLocation()

  return React.useMemo(() => new URLSearchParams(search), [search])
}

function setAuthJwt(jwt: string) {
  localStorage.setItem('jwt', jwt)
}

function resetAuthJwt() {
  localStorage.removeItem('jwt')
}

function AuthProvider(props: { children: React.ReactNode }) {
  const navigate = useNavigate()
  const query = useQuery()
  // check if is query param "jwt"
  // if so, set jwt in localstorage
  // and redirect to home
  useEffect(() => {
    const jwt = query.get('jwt')
    if (jwt) {
      setAuthJwt(jwt)
      navigate('/', { replace: true })
    }
  }, [query, navigate])

  return <>{props.children}</>
}

function useIsEmailConfirmed() {
  const tokenStr = localStorage.getItem('jwt')
  const jwtDecoded = jwtDecode(tokenStr || '') as { isVerified: boolean }
  return jwtDecoded?.isVerified === true
}

function useAuthInfo() {
  const tokenStr = localStorage.getItem('jwt')
  const jwtDecoded = jwtDecode(tokenStr || '') as { userId: boolean }
  return { userId: jwtDecoded?.userId }
}
export { resetAuthJwt, setAuthJwt, useIsEmailConfirmed, useAuthInfo }
export default AuthProvider
