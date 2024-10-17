import axios from "axios";

export const subscribeToNewsletter = async (email: string) => {
    try {
        const res = await axios.post(`${import.meta.env.VITE_BASE_URL}api/subscribeNewsletter`, {
            email
        })
        return res.data
    } catch (error) {
        return error
    }
}