import { useUpdatePassword } from '@/api/admin'
import { ApiError } from '@/api/gen'
import { useToast } from '@/components/Toasts/ToastsProvider'
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
  return <ChangePasswordModal />
}

function ChangePasswordModal() {
  const [open, setOpen] = useState(false)
  const update = useUpdatePassword()
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
            toast.error('Session expired - please log in again')
            return
          } else {
            toast.error('Error occurred - please try again')
          }
        },
      },
    )
    handleOpen()
  }

  return (
    <>
      <Button onClick={handleOpen} variant="gradient">
        Change password
      </Button>
      <form>
        <Dialog open={open} handler={handleOpen}>
          <DialogHeader>Password change</DialogHeader>
          <DialogBody className="flex flex-col pl-8 gap-4">
            <Input
              type="password"
              label="New password"
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
              Cancel
            </Button>
            <Button
              variant="gradient"
              color="green"
              onClick={handleConfirm}
              type="submit"
            >
              Change
            </Button>
          </DialogFooter>
        </Dialog>
      </form>
    </>
  )
}
export default Profile
