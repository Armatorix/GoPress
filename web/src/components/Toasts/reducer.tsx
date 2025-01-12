import { randomId } from '@/lib/random'
import type { ToastProps } from './Toast'
import type { Toasts } from './types'

export type Action =
  | { type: 'add'; body: ToastProps }
  | { type: 'remove'; body: { id: string } }

export function reducer(
  state: { toasts: Toasts[] },
  action: Action,
): { toasts: Toasts[] } {
  switch (action.type) {
    case 'add': {
      return { toasts: [...state.toasts, { ...action.body, id: randomId() }] }
    }
    case 'remove': {
      return {
        toasts: state.toasts.filter((toast) => toast.id !== action.body.id),
      }
    }
  }
}
