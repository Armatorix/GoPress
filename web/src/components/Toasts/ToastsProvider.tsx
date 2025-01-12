import {
  createContext,
  type PropsWithChildren,
  useContext,
  useReducer,
} from 'react'
import { reducer } from './reducer'
import { ToastWrapper } from './ToastWrapper'
import { ToastType, type ToastProps } from './Toast'

type ToastContextType = {
  success: (message: string) => void
  error: (message: string) => void
  info: (message: string) => void
}

export const ToastContext = createContext<ToastContextType | null>(null)

export function ToastProvider({ children }: PropsWithChildren) {
  const [{ toasts }, dispatchToasts] = useReducer(reducer, { toasts: [] })

  const createToast = (toast: ToastProps) => {
    dispatchToasts({ type: 'add', body: toast })
  }

  const removeToast = (id: string) => {
    dispatchToasts({ type: 'remove', body: { id: id } })
  }

  const success = (message: string) => {
    createToast({ message })
  }

  const error = (message: string) => {
    createToast({ message, type: ToastType.ERROR })
  }

  const info = (message: string) => {
    createToast({ message, type: ToastType.INFO })
  }

  return (
    <ToastContext.Provider value={{ success, error, info }}>
      <>
        {children}
        {toasts.length > 0 && (
          <div className="fixed bottom-0 right-16 items-center flex w-full flex-col  gap-2 py-6">
            {toasts.map(({ id, ...rest }) => (
              <ToastWrapper
                key={id}
                dismiss={() => removeToast(id)}
                {...rest}
              />
            ))}
          </div>
        )}
      </>
    </ToastContext.Provider>
  )
}

export function useToast() {
  const context = useContext(ToastContext)

  if (!context) {
    throw new Error(
      'Child components of Toasts cannot be rendered outside the Toast provider!',
    )
  }

  return context
}
