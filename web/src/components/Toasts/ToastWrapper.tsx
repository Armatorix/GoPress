import { useEffect } from 'react'
import { Toast, type ToastProps } from './Toast'

type ToastWrapperProps = ToastProps & { dismiss: () => void }

export function ToastWrapper({ dismiss, ...rest }: ToastWrapperProps) {
  useEffect(() => {
    const id = setTimeout(dismiss, 2000)

    return () => clearTimeout(id)
  }, [])

  return <Toast {...rest} />
}
