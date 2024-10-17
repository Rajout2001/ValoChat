import React from 'react'

interface InputProps {
    placeholder: string;
    stateValue: string | undefined;
    setState: (stateName: string) => void;
    type?: string;
}

const Input = (inputProps : InputProps) => {
    const { placeholder, stateValue, setState , type} = inputProps;
  return (
    <input
        className="px-4 py-2 bg-red-500 text-white rounded-lg focus:outline-none focus:ring-2 focus:ring-red-600 focus:ring-opacity-100 placeholder-white"
        type={type ? type : "text"}
        placeholder={placeholder}
        value={stateValue}
        onChange={(e) => setState(e.target.value)}
    />
  )
}

export default Input