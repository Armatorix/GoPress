import { useLogin } from '@/api/auth'
import type { AuthToken, PostLogin } from '@/api/gen'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { setAuthJwt } from '@/lib/AuthProvider'
import Fullpage from '@/page_wrappers/Fullpage'
import { Button, Input } from '@material-tailwind/react'
import { useState } from 'react'

function Login() {
  const login = useLogin()
  const toast = useToast()
  const [authData, setAuthData] = useState<PostLogin>({
    email: '',
    password: '',
  })

  const onPasswordLoginClick = (e: any) => {
    e.preventDefault()
    login.mutate(authData, {
      onSuccess: (data: AuthToken) => {
        setAuthJwt(data.token)
        window.location.reload()
      },
      onError: (_: any) => {
        toast.error('Niepoprawny login lub hasło')
      },
    })
  }

  return (
    <Fullpage
      className="md:w-[24em] flex flex-col gap-4"
      title={
        <div className="flex flex-row gap-2 items-center">
          <img src="/icon.png" alt="logo" className="w-10 h-10" />
          GoPress
        </div>
      }
    >
      <form
        onSubmit={onPasswordLoginClick}
        className="text-center flex flex-col gap-2"
      >
        <Input
          type="email"
          label="Email"
          value={authData.email}
          onChange={(e) => setAuthData({ ...authData, email: e.target.value })}
        />
        <Input
          type="password"
          label="Password"
          value={authData.password}
          onChange={(e) =>
            setAuthData({ ...authData, password: e.target.value })
          }
        />
        <Button
          variant="outlined"
          onClick={() => onPasswordLoginClick}
          type="submit"
        >
          Login
        </Button>
      </form>
    </Fullpage>
  )
}

export default Login
