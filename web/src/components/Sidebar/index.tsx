import { resetAuthJwt } from '@/lib/AuthProvider'
import {
  ClipboardDocumentIcon,
  Cog6ToothIcon,
  PowerIcon,
  PresentationChartBarIcon,
} from '@heroicons/react/24/solid'
import {
  Card,
  List,
  ListItem,
  ListItemPrefix,
  Typography,
} from '@material-tailwind/react'

function Sidebar() {
  const onLogout = () => {
    resetAuthJwt()
    window.location.reload()
  }
  return (
    <Card className="h-screen w-full max-w-72 p-4 shadow-xl shadow-blue-gray-900">
      <div className="mb-2 p-4 flex flex-row items-center justify-center space-x-2 select-none">
        <img src="/icon.png" alt="logo" className="h-10 w-10" />
        <Typography variant="h5" color="blue-gray">
          GoPress
        </Typography>
      </div>
      <List>
        <ListItem disabled>
          <ListItemPrefix>
            <PresentationChartBarIcon className="h-5 w-5" />
          </ListItemPrefix>
          Dashboard
        </ListItem>
        <ListItem disabled>
          <ListItemPrefix>
            <ClipboardDocumentIcon className="h-5 w-5" />
          </ListItemPrefix>
          Articles
        </ListItem>
        <ListItem disabled>
          <ListItemPrefix>
            <Cog6ToothIcon className="h-5 w-5" />
          </ListItemPrefix>
          Settings
        </ListItem>
        <ListItem onClick={onLogout}>
          <ListItemPrefix>
            <PowerIcon className="h-5 w-5" />
          </ListItemPrefix>
          Log Out
        </ListItem>
      </List>
    </Card>
  )
}

export default Sidebar
