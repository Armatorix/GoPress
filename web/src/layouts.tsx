import Navbar from '@/components/Navbar'
import { Outlet } from 'react-router-dom'

function AuthedMainLayout() {
  return (
    <>
      <Navbar />
      <MainLayout>
        <Outlet />
      </MainLayout>
    </>
  )
}

function UnauthedMainLayout() {
  return (
    <div className="w-full h-full">
      <MainLayout>
        <Outlet />
      </MainLayout>
    </div>
  )
}
function MainLayout(props: { children: React.ReactNode }) {
  return <div className="w-full h-full">{props.children}</div>
}

export { AuthedMainLayout, UnauthedMainLayout }
