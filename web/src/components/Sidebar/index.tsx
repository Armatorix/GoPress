import { resetAuthJwt } from '@/lib/AuthProvider'
import {
  Bars3Icon,
  ClipboardDocumentIcon,
  Cog6ToothIcon,
  PowerIcon,
  PresentationChartBarIcon,
  XMarkIcon,
} from '@heroicons/react/24/solid'
import {
  Card,
  Drawer,
  IconButton,
  List,
  ListItem,
  ListItemPrefix,
  Typography,
} from '@material-tailwind/react'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

function Sidebar() {
  return (
    <>
      <div className="hidden xl:block">
        <AlwaysOpenSidebar />
      </div>
      <div className="xl:hidden">
        <DrawerSidebar />
      </div>
    </>
  )
}

function AlwaysOpenSidebar() {
  return (
    <Card className="h-screen fixed">
      <SidebarContent />
    </Card>
  )
}

function DrawerSidebar() {
  const [open, setOpen] = useState(false)
  return (
    <>
      <IconButton variant="text" size="lg" onClick={() => setOpen(!open)}>
        {open ? (
          <XMarkIcon className="h-8 w-8 stroke-2" />
        ) : (
          <Bars3Icon className="h-8 w-8 stroke-2" />
        )}
      </IconButton>

      <Drawer open={open} onClose={() => setOpen(false)}>
        <SidebarContent />
      </Drawer>
    </>
  )
}

function SidebarContent() {
  const navigate = useNavigate()

  const onDashboard = () => {
    navigate('/dashboard')
  }

  const onArticles = () => {
    navigate('/articles')
  }
  const onLogout = () => {
    resetAuthJwt()
    window.location.reload()
  }

  const onProfile = () => {
    navigate('/profile')
  }
  return (
    <Card className="h-screen">
      <div className="mb-2 p-4 flex flex-row items-center justify-center space-x-2 select-none">
        <img src="/icon.png" alt="logo" className="h-10 w-10" />
        <Typography variant="h5" color="blue-gray">
          GoPress
        </Typography>
      </div>
      <List className='min-w-48'>
        <ListItem onClick={onDashboard}>
          <ListItemPrefix>
            <PresentationChartBarIcon className="h-5 w-5" />
          </ListItemPrefix>
          Dashboard
        </ListItem>
        <ListItem onClick={onArticles}>
          <ListItemPrefix>
            <ClipboardDocumentIcon className="h-5 w-5" />
          </ListItemPrefix>
          Articles
        </ListItem>
        <ListItem onClick={onProfile}>
          <ListItemPrefix>
            <Cog6ToothIcon className="h-5 w-5" />
          </ListItemPrefix>
          Profile
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
