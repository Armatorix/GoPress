import { useUser } from '@/api'
import { ApiError } from '@/api/gen'
import { updatePassword } from '@/api/user'
import { useToast } from '@/components/Toasts/ToastsProvider'
import { Bootstrap } from '@/page_wrappers/Bootstrap'
import {
  Button,
  Dialog,
  DialogBody,
  DialogFooter,
  DialogHeader,
  Input,
} from '@material-tailwind/react'
import { useState } from 'react'

function Profile() {
  const user = useUser()
  const userData = user.data?.data
  const { profile } = useOrganizationPaths()
  if (!user || user.isLoading || !userData) {
    return null
  }
  return (
    <Bootstrap>
      <div className="flex flex-col gap-4">
        <div className="flex flex-row gap-4">
          {userData.avatarURL && (
            <div>
              <img
                src={userData.avatarURL}
                alt={`${userData.firstName.at(0)} ${userData.lastName.at(0)}`}
                className="rounded-full w-24 h-24 text-center"
              />
            </div>
          )}
          <div className="flex flex-col gap-2">
            <Input
              type="email"
              label="Imię"
              variant="outlined"
              disabled
              value={userData.firstName}
            />
            <Input
              type="text"
              label="Nazwisko"
              variant="outlined"
              disabled
              value={userData.lastName}
            />
          </div>
        </div>
        <ChangePasswordModal />
      </div>
    </Bootstrap>
  )
}

function ChangePasswordModal() {
  const [open, setOpen] = useState(false)
  const update = updatePassword()
  const toast = useToast()
  const handleOpen = () => setOpen(!open)
  const [password, setPassword] = useState('')
  const handleConfirm = () => {
    update.mutate(
      { password },
      {
        onSuccess: () => {
          toast.success('Hasło zostało zmienione')
        },
        onError: (error) => {
          // check if error is type of ApiError
          if (error instanceof ApiError && error.status === 412) {
            toast.error(
              'Sesja trwa zbyt długo - zaloguj się ponownie i spróbuj ponownie',
            )
            return
          } else {
            toast.error('Wystąpił błąd - spróbuj ponownie')
          }
        },
      },
    )
    handleOpen()
  }

  return (
    <>
      <Button onClick={handleOpen} variant="gradient">
        Zmień hasło
      </Button>
      <form>
        <Dialog open={open} handler={handleOpen}>
          <DialogHeader>Zmiana hasła</DialogHeader>
          <DialogBody className="flex flex-col pl-8 gap-4">
            <Input
              type="password"
              label="Nowe hasło"
              variant="outlined"
              value={password}
              required
              onChange={(e) => setPassword(e.target.value)}
            />
          </DialogBody>
          <DialogFooter>
            <Button
              variant="text"
              color="red"
              onClick={handleOpen}
              className="mr-1"
            >
              Anuluj
            </Button>
            <Button
              variant="gradient"
              color="green"
              onClick={handleConfirm}
              type="submit"
            >
              Zmień
            </Button>
          </DialogFooter>
        </Dialog>
      </form>
    </>
  )
}
export default Profile
