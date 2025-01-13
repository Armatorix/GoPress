import Sidebar from '@/components/Sidebar'
import clsx from 'clsx'

type BootstrapProps = {
  children: React.ReactNode
  className?: string
}
function Bootstrap({ children, className }: BootstrapProps) {
  return (
    <>
      <Sidebar />
      <div className="flex flex-col gap-2 pt-24 pb-14 max-w-[65em] m-auto w-[95%]">
        <div
          className={clsx(
            'border shadow p-10 flex flex-col gap-4 bg-white rounded-xl z-0 m-auto w-full',
            className,
          )}
        >
          <div className="w-full"> {children} </div>
        </div>
      </div>
    </>
  )
}

export { Bootstrap }
