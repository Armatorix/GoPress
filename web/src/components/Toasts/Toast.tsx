import clsx from 'clsx'

export enum ToastType {
  DEFAULT = 'default',
  ERROR = 'error',
  INFO = 'info',
}
export type ToastProps = {
  message: string
  type?: ToastType
}

export function Toast({ message, type = ToastType.DEFAULT }: ToastProps) {
  return (
    <section
      className={clsx(
        'w-fit rounded-lg px-4 py-3 shadow',
        type === ToastType.DEFAULT && 'bg-green-500 text-white',
        type === ToastType.ERROR && 'bg-red-500 text-dusk',
        type === ToastType.INFO && 'bg-blue-500 text-white',
      )}
      aria-label={`toast: ${message}`}
    >
      <div>{message}</div>
    </section>
  )
}
