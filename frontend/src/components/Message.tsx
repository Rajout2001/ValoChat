import { Message } from "../services/ChatService"

// interface MessageProps {
//     content: string
//     user_id: string
// }

const MessageBox = ({user, content}: Message) => {
  return (
    <div className='flex flex-row justify-start items-center gap-2 h-fit m-2'>
        <div>
            {user.name}:
        </div>
        <div className='p-2 bg-gray-600 rounded-lg'>
            {content}
        </div>
    </div>
  )
}

export default MessageBox