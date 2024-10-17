import { useNavigate } from "react-router-dom";
import { FormEvent, useEffect, useState } from "react";
import { login, register } from "../services/AuthService";
import { useAuth } from "../context/AuthProvider";
import { validateEmail, validatePassword } from "../utils";
import LoginBox from "../components/LoginBox";
import RegisterBox from "../components/RegisterBox";

interface Error {
    message: string;
    code : string;
}

const Login = () => {
    const navigate = useNavigate();
    const { user } = useAuth();
    const [email, setEmail] = useState('hichem.bouzourine@etu.sorbonne-universite.fr');
    const [password, setPassword] = useState('projet_pc3r');
    const [error, setError] = useState<Error>()
    const [Login, setLogin] = useState<boolean>(true);
    const [name, setName] = useState<string>('')

    useEffect(() => {
        // Check if user is logged in after successful login
        if (user) {
            navigate('/');
        }
    }, [user, navigate]);

    const handleLoginSubmit = async (e: FormEvent) => {
        e.preventDefault();
        // regex for email validation
        if (!validateEmail(email)){
            return;
        }

        if (!validatePassword(password)){
            return;
        }
        
        const data = await login(email, password);
        if (data?.tokens) {
            window.location.href = '/';
        }else {
            setError(data.response.data.error)
        }
    }

    const handleRegisterSubmit = async (e: FormEvent) => {
        e.preventDefault();
        if (!name) {
            alert('Please enter your name');
            return;
        }
        if (!password) {
            alert('Please enter a password');
            return;
        }
        
        if (!validateEmail(email)){
            return;
        }

        if (!validatePassword(password)){
            return;
        }
        
        const data = await register(email, password, name);
        if (data?.tokens) {
            window.location.href = '/';
        }else {
            setError(data.response.data.error)
        }
    }
    
    return (
        <div className="flex flex-col items-center justify-center h-screen bg-gray-800">
            <div className="w-1/5">
                <img src={"/valorant.png"} alt="" />
            </div>
            <div className="flex flex-col items-center justify-center w-2/5 h-2/5 shadow-slate-400">
                    {Login && (
                        <LoginBox 
                            email={email}
                            setEmail={setEmail}
                            password={password}
                            setPassword={setPassword}
                            error={error}
                            setError={setError}
                            setLogin={setLogin}
                            handleLoginSubmit={handleLoginSubmit}
                        />
                    )}
                    {!Login && (
                        <RegisterBox 
                            email={email}
                            setEmail={setEmail}
                            password={password}
                            setPassword={setPassword}
                            error={error}
                            setError={setError}
                            setLogin={setLogin}
                            handleRegisterSubmit={handleRegisterSubmit}
                            name={name}
                            setName={setName}
                        />
                    )}
            </div>
        </div>
    );
};

export default Login;