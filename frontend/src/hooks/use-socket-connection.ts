import { useEffect } from "react";
import { useWebSocket } from "../context/WebSocketProvider";
import { useAuth } from "../context/AuthProvider";



export const useSocketConnection = ()=>{
    const {socket, setSocket} = useWebSocket();
    const {user} = useAuth();

  useEffect(() => {
    if (socket) return

    if (!localStorage.getItem("token")) return
    const newSocket = new WebSocket(`${import.meta.env.VITE_BACKEND_WS_URL}?Authorization=Bearer ${localStorage.getItem("token")}`);
    setSocket(newSocket);
    return ()=>{
        console.log("trying to disconnect ! ",newSocket?.readyState)
        if (newSocket?.readyState  === 1) {
            // notifier le serveur que le client se détache 
            const message = JSON.stringify({Event:"unsubscribe", Data:"Closing"});
            newSocket.send(message);
        }

    }

  }, [user]);

}