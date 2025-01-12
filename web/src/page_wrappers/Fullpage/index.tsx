import { Typography } from '@material-tailwind/react'
import clsx from 'clsx'

type FullpageProps = {
  title: any
  children: React.ReactNode
  className?: string
}
export function Fullpage({ title, children, className }: FullpageProps) {
  return (
    <div className="flex justify-center py-64">
      <div
        className={clsx(
          'border px-16 py-10 flex flex-col justify-center items-center gap-4 bg-white rounded-xl z-0 max-w-[65em] m-auto w-[95%]',
          className,
        )}
      >
        <Typography color="gray" variant="h4">
          {title}
        </Typography>
        <>{children}</>
      </div>
    </div>
  )
}

export default Fullpage
