'use server'

interface LinkResponse {
  success: boolean
  message: string
  shortLink?: string
}

export const generateShortLink = async (prevState: LinkResponse, formData: FormData) => {
  const url = formData.get('url') as string

  try {
    if (!url) {
      return {
        success: false,
        message: 'URL is required'
      }
    }

    const payload = { url: url }
    const apiUrl = process.env.NEXT_PUBLIC_API_URL
    const response = await fetch(`${apiUrl}/shorten`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload),
    })

    console.log('Response status:', response.status)
    const text = await response.text()
    console.log('Response text:', text)

    let data
    try {
      data = JSON.parse(text)
    } catch (error) {
      console.error("Failed to parse JSON:", error)
      return {
        success: false,
        message: 'Invalid JSON response'
      }
    }

    if (!response.ok) {
      throw new Error(data.message || 'Unknown error')
    }

    const baseApiUrl = process.env.NEXT_PUBLIC_BASE_API_URL
    const fullShortLink = `${baseApiUrl}/${data.shortLink}`

    return {
      success: true,
      message: 'Short link generated successfully',
      shortLink: fullShortLink
    }
  } catch (error) {
    console.log(error)
    if (error instanceof Error) {
      return {
        success: false,
        message: error.message
      }
    }
    return {
      success: false,
      message: 'Failed to generate short link'
    }
  }
}