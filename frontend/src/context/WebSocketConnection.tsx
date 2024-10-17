import { ReactNode } from 'react';
import { useSocketConnection } from '../hooks/use-socket-connection';

interface WebSocketConnectionProps {
    children : ReactNode
}

function WebSocketConnection({children}:WebSocketConnectionProps) {
    useSocketConnection()
    return (
    <>        
    {children}
    </>

    );
}

export default WebSocketConnection;