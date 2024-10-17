import React, { ButtonHTMLAttributes } from 'react'

interface ButtonProps {
    type?: ButtonHTMLAttributes<HTMLButtonElement>['type'];
    text : string;
}

const Button = (buttonProps : ButtonProps) => {
    const { type, text } = buttonProps;
  return (
    <>
        <button 
            type={type ? type : "submit"}
            className="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-400 focus:outline-none focus:ring-2 focus:ring-red-600 focus:ring-opacity-50 hover:cursor-pointer"
        >
            {text}
        </button>
    </>
  )
}

export default Button