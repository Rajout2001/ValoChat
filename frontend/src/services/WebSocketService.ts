import { useEffect, useState } from "react";

const useWebSocket = () => {
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [socketIsOpen, setSocketIsOpen] = useState<boolean>(false);
  const [error, setError] = useState<Event | null>(null);

  useEffect(() => {
    if (socket) {
      console.log("Socket already exists");
      return
    }

    if (!localStorage.getItem("token")) {
      console.log("No token found");
      return
    }

    const newSocket = new WebSocket(`${import.meta.env.VITE_BACKEND_WS_URL}?Authorization=Bearer ${localStorage.getItem("token")}`);
    setSocket(newSocket);
    console.log("Socket created", newSocket);

  }, [localStorage.getItem("token")]);


  useEffect(() => {
    if (!socket) return;

    socket.onopen = () => {
      console.log("WebSocket connected");
      setSocketIsOpen(true);
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
      setError(error);
    };

    socket.onclose = () => {
      console.log("WebSocket disconnected");
      sendEvent("unsubscribe", "unsubscribe");
      setSocketIsOpen(false);
    };

    return () => {
      console.log("Socket is closing");
      if (socket) {
        socket.onopen = null;
        socket.onerror = null;
        socket.onclose = null;
        setSocket(null);
        socket.close();
        console.log("WebSocket disconnected from useEffect");
      }
    }
  }, [localStorage.getItem("token"), socket]);

  const sendEvent = (event : string, data : any) => {
    if (socket && socketIsOpen) {
      socket.send(JSON.stringify({ event, data}));
    } else {
      console.error(`${event} not send, socket not connected.`);
    }
  };

  return { socket, socketIsOpen, error, sendEvent };
};

export default useWebSocket;