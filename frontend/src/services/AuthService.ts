import axios from "axios";

export const login = async (email: string, password: string) => {
    try {
        const res = await axios.post(`${import.meta.env.VITE_BASE_URL}api/auth/login`, {
            email,
            password
        })
        const data = res.data
        localStorage.setItem('connectedUser', JSON.stringify(data?.user));
        localStorage.setItem('token', data?.tokens?.access);
        return res.data;
    } catch (error) {
        console.error(error)
        return error;
    }
    
};

export const register = async (email: string, password: string, name: string) => {
    try {
        const res = await axios.post(`${import.meta.env.VITE_BASE_URL}api/auth/signup`, {
            email,
            password,
            name
        })
        const data = res.data
        localStorage.setItem('connectedUser', JSON.stringify(data?.user));
        localStorage.setItem('token', data?.tokens?.access);        
        
        return data;
    } catch (error) {
        console.error(error)
        return error;
    }
}