import axios from "axios"

export interface Chat {
    id: string
    date: string
    name: string
    photo: string
    messages: Message[]
}

export interface Message {
    id?: string
    content : string
    chat_id?: string
    user_id?: string
    user : User
}

export interface User {
    id : string
    name: string
    photo: string
}

interface getChatsProps {
    chatId: string
    setChat: React.Dispatch<React.SetStateAction<Chat | null>> 
}

export const getChat = async ({chatId, setChat}: getChatsProps) => {
    const url = `${import.meta.env.VITE_BASE_URL}api/chat?id=${chatId}`
    const res = await axios.get(url, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`
        }
    })
    const {chat} = res.data
    setChat(chat)
}