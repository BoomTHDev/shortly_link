'use client'

import { generateShortLink } from '@/actions/link'
import Form from 'next/form'
import { useActionState } from 'react'

const NewLinkForm = () => {
  const prevState = {
    success: false,
    message: '',
    shortLink: '',
  }
  const [state, action, isPending] = useActionState(
    generateShortLink,
    prevState
  )

  return (
    <div className='flex flex-col gap-4'>
      <Form
        action={action}
        className='flex gap-4 items-center'
      >
        <input
          type='url'
          name='url'
          placeholder='Enter your URL'
          className='border border-gray-300 rounded-lg p-2 w-full'
        />
        <button
          type='submit'
          disabled={isPending}
          className='bg-indigo-600 text-white px-4 py-2 rounded-lg'
        >
          {isPending ? 'Generating...' : 'Shorten'}
        </button>
      </Form>
      {state.success ? (
        <p className='text-green-500'>
          {state.message}
          <br />
          <a
            href={state.shortLink}
            target='_blank'
            rel='noopener noreferrer'
          >
            {state.shortLink}
          </a>
        </p>
      ) : (
        <p className='text-red-500'>{state.message}</p>
      )}
    </div>
  )
}
export default NewLinkForm
