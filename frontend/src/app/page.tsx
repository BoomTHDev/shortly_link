import NewLinkForm from '@/components/new_link_form'

const HomePage = () => {
  return (
    <div className='min-h-svh bg-gradient-to-br from-blue-950 to-indigo-900 py-12 px-4 sm:px-6 lg:px-8'>
      <div className='max-w-3xl mx-auto'>
        <div className='text-center mb-12'>
          <h1 className='text-4xl font-extrabold text-gray-100 sm:text-5xl md:text-6xl mb-4'>
            <span className='block'>Shorten Your</span>
            <span className='block text-indigo-600'>Links in Seconds</span>
          </h1>
          <p className='mt-3 max-w-md mx-auto text-base text-gray-300 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl'>
            Create short, memorable links that are perfect for sharing.
          </p>
        </div>
        <div className='bg-gray-800 rounded-2xl shadow-xl p-6 sm:p-8'>
          <NewLinkForm />
        </div>
      </div>
    </div>
  )
}

export default HomePage
