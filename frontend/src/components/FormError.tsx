import React from 'react'

interface FormErrorProps {
    error: any;
}

const FormError = (formError : FormErrorProps) => {
    const { error } = formError;
  return (
    <div className="border p-3 my-3 border-red-500">
        <p className="text-red-500">Error: {error?.message}</p>
    </div>
  )
}

export default FormError