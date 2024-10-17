import { FormEvent, useEffect, useRef, useState } from "react";
import { Chat, Message } from "../services/ChatService";
import MessageBox from "./Message";
import { useWebSocket } from "../context/WebSocketProvider";

interface ChatBoxProps {
    chat: Chat | null
}

const ChatBox = ({chat}: ChatBoxProps) => {
    const chatBoxRef = useRef<HTMLDivElement>(null);
    const [messageInput, setMessageInput] = useState<string>('');
    const [messages, setMessages] = useState<Message[] | undefined>(chat?.messages);
    const {socket, socketIsOpen, error, sendEvent} = useWebSocket()
   
 
    useEffect(() => {
        // If there is an websocket error, return
        if (error) return;

    }, [socketIsOpen, chat]);

    useEffect(() => {
      // if there is no socket, return
      if (!socket) return;

      // Listen for messages from the server
      socket.onmessage = (event) => {
        const message = JSON.parse(event.data);
        if (message.event === "receive_message") {
          setMessages((prevMessages = []) => [...prevMessages, message.data]);
        }
      }
    } , [socket]);

    
    useEffect(() => {
        // Update the messages when the chat updates
        setMessages(chat?.messages);
        // Scroll to the bottom of the chat box when component mounts or chat updates
        if (chatBoxRef.current) {
            chatBoxRef.current.scrollTop = chatBoxRef.current.scrollHeight;
        }
    }, [chat]);

    const sendMessage = (e: FormEvent) => {
        e.preventDefault();  
        // If the socket is open and there is a chat and a message input send the message
        if (socketIsOpen && chat && messageInput && (sendEvent !== null)) {
          sendEvent("send_message", { chat_id: chat?.id, content: messageInput });
        }
      
        setMessageInput("");
    };

    return (
        <div className="flex flex-col justify-center">
            <div className="flex flex-row justify-between gap-5 p-2 border-b">
                <div className="flex justify-center items-center orange_gradient">{chat?.name}</div>
                <div className="w-12">
                    <img src={chat?.photo} alt="chat img" />
                </div>
            </div>
            <div ref={chatBoxRef} className="ml-2 flex flex-col-reverse border-b h-[500px] overflow-auto">
                {/* <div className="max-h-[500px]"> */}
                    {
                        messages?.slice().reverse().map((message, index) => (
                            <MessageBox key={index} user={message.user} content={message.content} />
                        ))
                    }
                {/* </div> */}
            </div>
            <form className="p-2 my-2 flex flex-row justify-between gap-5 rounded-lg" onSubmit={sendMessage}>
                <input type="text" placeholder="Type a message" className="w-full rounded-lg p-2 text-slate-700"
                    value={messageInput}
                    onChange={(e) => setMessageInput(e.target.value)} 
                />
                <button 
                    type="submit"
                    className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                    // onClick={() => {sendMessage()}}
                >
                    Send
                </button>
            </form>
        </div>
    )
}

export default ChatBox;