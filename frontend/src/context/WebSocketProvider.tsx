import { createContext, useContext, useEffect, useMemo, useState } from "react";

const WebSocketContext = createContext({
  socket: null as WebSocket | null,
  socketIsOpen: false,
  setSocket: (socket:WebSocket | null)=>{},
  error: null as Event | null,
  sendEvent: (event: string, data: any) => {}
});

export const useWebSocket = () => useContext(WebSocketContext);

export const WebSocketProvider = ({ children }: { children?: React.ReactNode }) => {
  const memToken = useMemo(()=>localStorage.getItem('token'),[localStorage.getItem('token')])
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [socketIsOpen, setSocketIsOpen] = useState<boolean>(false);
  const [error, setError] = useState<Event | null>(null);

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
      setSocketIsOpen(false);
      setSocket(null)
      socket.close();
    };

  } , [socket]);

  const sendEvent = (event: string, data: any) => {
    if (socket && socketIsOpen) {
      socket.send(JSON.stringify({ event, data }));
    } else {
      console.error(`${event} not sent, socket not connected.`);
    }
  };

  return (
    <WebSocketContext.Provider
      value={{ socket, socketIsOpen, error, sendEvent, setSocket: (setSocket as any) }}
    >
      {children}
    </WebSocketContext.Provider>
  );
};
