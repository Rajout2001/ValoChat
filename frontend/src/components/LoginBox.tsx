import React from 'react'
import Input from './Input';
import Button from './Button';
import FormError from './FormError';

interface LoginBoxProps {
    email?: string;
    setEmail: (email: string) => void;
    password?: string;
    setPassword: (password: string) => void;
    error: any;
    setError: (error: any) => void;
    setLogin: (login: boolean) => void;
    handleLoginSubmit?: (e: React.FormEvent) => void;
}

const LoginBox = (loginProps: LoginBoxProps) => {
    const { email, setEmail, password, setPassword, error, setError, setLogin, handleLoginSubmit } = loginProps;
  return (
    <>
        <h1 className="text-3xl font-bold text-red-500 mb-5">Login</h1>
        <form className="flex flex-col items-center justify-center gap-y-4" onSubmit={handleLoginSubmit}>
            <Input placeholder="Email" stateValue={email} setState={setEmail} />
            <Input placeholder="Password" stateValue={password} setState={setPassword} type='password'/>
            <Button text="Login" />
            {error && (
                <FormError error={error} />   
            )}
            <div>
                <p className="text-white">Don't have an account? <span onClick={() => {setLogin(false); setError(undefined)}} className="text-red-500 hover:cursor-pointer">Register</span></p>
            </div>
        </form>
    </>
  )
}

export default LoginBox