package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pc3r/prisma"
	"pc3r/prisma/db"
	"strings"
)



type ChatRes struct {
	*db.ChatModel
	Users     []UserChatModel       `json:"users"`
	Messages  []MessageChatResponse `json:"messages"`
	Messagess []*db.MessageModel    `json:"Messages"`
}

type ResponseGetChatBody struct {
	Chat ChatRes `json:"chat"`
}

func GetChat(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		res.WriteHeader(http.StatusUnauthorized)
		message := "Unauthorized"
		json.NewEncoder(res).Encode(CustomError(message, UNAUTHORIZED))
		return
	}
	
	// user, _ := req.Context().Value(CtxAuthKey{}).(*db.UserModel)
	prisma, ctx := prisma.GetPrisma()
	chat, err := prisma.Chat.FindFirst(
		db.Chat.ID.Equals(id),
	).With(
		db.Chat.Users.Fetch(),
		db.Chat.Messages.Fetch().With(
			db.Message.User.Fetch(),
		),
		db.Chat.Match.Fetch(),
	).Exec(ctx)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		message := "Chat with ID '" + id + "'doesn't exist." 
		json.NewEncoder(res).Encode(CustomError(message, NOT_FOUND))
		return
	}

	users := ExtractChatUsersInformations(chat.Users())
	messages := StructureMessages(chat.Messages())
	chatStructure := ChatRes{
		Messages:  messages,
		ChatModel: chat,
		Users:     users,
	}
	response := ResponseGetChatBody{
		Chat: chatStructure,
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)

}

type CreateChatProps struct {
	Id string `json:"id"`
}


type UserChatModel struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo string `json:"photo"`
	Id    string `json:"id"`
}
func ExtractChatUsersInformations(chat_users []db.UserModel) []UserChatModel {
	users := []UserChatModel{}
	for _, user := range chat_users {
		photo, _ := user.Photo()
		// extract only necessary informations such as : id, name , email , photo
		updated_user := UserChatModel{
			Name:  user.Name,
			Email: user.Email,
			Photo: photo,
			Id:    user.ID,
		}
		users = append(users, updated_user)
	}
	return users
}

func StructureMessages(messages []db.MessageModel) []MessageChatResponse {
	structured_messages := []MessageChatResponse{}
	for _, message := range messages {
		structured_message := StructureMessage(message)
		structured_messages = append(structured_messages, structured_message)
	}
	return structured_messages
}

type MessageChatResponse struct {
	Id         string        `json:"id"`
	Content    string        `json:"content"`
	Chat_id    string        `json:"chat_id"`
	User_id    string        `json:"user_id"`
	Created_at db.DateTime   `json:"created_at"`
	User       UserChatModel `json:"user"`
}

func StructureMessage(message db.MessageModel) MessageChatResponse {
	m_user := message.User()
	photo, _ := m_user.Photo()

	message_user := UserChatModel{
		Id:    m_user.ID,
		Name:  m_user.Name,
		Photo: photo,
		Email: m_user.Email,
	}
	structured_message := MessageChatResponse{
		Id:         message.ID,
		Content:    message.Content,
		Chat_id:    message.ChatID,
		User_id:    message.UserID,
		Created_at: message.CreatedAt,
		User:       message_user,
	}
	return structured_message
}

/* added user */
type ResponseAddUserToChatBody struct {
	Chat ChatRes `json:"chat"`
}

func AddUserToChat(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		res.WriteHeader(http.StatusUnauthorized)
		message := "Unauthorized"
		json.NewEncoder(res).Encode(CustomError(message, UNAUTHORIZED))
		return
	}

	prisma, ctx := prisma.GetPrisma()
	// retrieve the requested chat
	_, err := prisma.Chat.FindFirst(
		db.Chat.ID.Equals(id),

	).Exec(ctx)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		message := "Chat with ID: " + id + " Not found"
		json.NewEncoder(res).Encode(CustomError(message, NOT_FOUND))
		return
	}
	user, _ := req.Context().Value(CtxAuthKey{}).(*db.UserModel)

	// update the chat by adding a user
	updated_chat, err := prisma.Chat.FindUnique(
		db.Chat.ID.Equals(id),
	).With(
		db.Chat.Match.Fetch(),
		db.Chat.Users.Fetch(),
		db.Chat.Messages.Fetch(),
	).Update(
		db.Chat.Users.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(ctx)
	if err != nil {
		fmt.Printf("my error is : %s " , err)
		res.WriteHeader(http.StatusNotFound)
		message := "Chat Not found"
		json.NewEncoder(res).Encode(CustomError(message, NOT_FOUND))
		return
	}

	chat_users := updated_chat.Users()
	users := ExtractChatUsersInformations(chat_users)
	messages := StructureMessages(updated_chat.Messages())

	chatStructure := ChatRes{
		Messages:  messages,
		ChatModel: updated_chat,
		Users:     users,
	}
	response := ResponseAddUserToChatBody{
		Chat: chatStructure,
	}
	// ! SEND SOCKET TO OTHERS
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
}

func RemoveUserFromChat(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		res.WriteHeader(http.StatusUnauthorized)
		message := "Unauthorized"
		json.NewEncoder(res).Encode(CustomError(message, NOT_FOUND))
		return
	}

	prisma, ctx := prisma.GetPrisma()
	// retrieve current user from the request
	user, _ := req.Context().Value(CtxAuthKey{}).(*db.UserModel)

	_, err := prisma.Chat.FindFirst(
		db.Chat.ID.Equals(id),
		db.Chat.Users.Some(
			db.User.ID.Equals(user.ID),
		),
	).Exec(ctx)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		message := "You are not member of this Chat."
		json.NewEncoder(res).Encode(CustomError(message, NOT_FOUND))
		return
	}
	_, err = prisma.Chat.FindUnique(
		db.Chat.ID.Equals(id),
	).With(
		db.Chat.Messages.Fetch(),
		db.Chat.Match.Fetch(),
		db.Chat.Users.Fetch(),
	).Update(
		db.Chat.Users.Unlink( // removing the current user from the chat
			db.User.ID.Equals(user.ID),
		),
	).Exec(ctx)

	if err != nil {
		message := "Error while trying to remove current user from the chat"
		json.NewEncoder(res).Encode(CustomError(message, BAD_REQUEST))
		return
	}

	//! SEND SOCKET TO OTHER USERS
	res.WriteHeader(http.StatusAccepted)
	message := "Left Channel of " + id
	json.NewEncoder(res).Encode(MessageResponse{Message: message})
}

type SendMessageChatProps struct {
	Message string `json:"message"`
}

type ResponseSendMessage struct {
	Message *db.MessageModel `json:"message"`
}

func SendMessageService(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	var body SendMessageChatProps
	err := json.NewDecoder(req.Body).Decode(&body)
	if (err != nil) || strings.TrimSpace(body.Message) == "" {
		res.WriteHeader(http.StatusUnauthorized)
		message := "Missing properties"
		json.NewEncoder(res).Encode(CustomError(message, INPUT_ERROR))
		return
	}
	if id == "" {
		res.WriteHeader(http.StatusUnauthorized)
		message := "Unauthorized"
		json.NewEncoder(res).Encode(CustomError(message, UNAUTHORIZED))
		return
	}

	user, _ := req.Context().Value(CtxAuthKey{}).(*db.UserModel)
	prisma, ctx := prisma.GetPrisma()
	// Get a chat also fetch the users and messages
	chat, err := prisma.Chat.FindFirst(
		db.Chat.ID.Equals("clv5q7mud0000of7bprqa5u41"),
		db.Chat.Users.Some(
			db.User.ID.Equals(user.ID),
		),
	).With(
		db.Chat.Users.Fetch(),
		db.Chat.Messages.Fetch().With(
			db.Message.User.Fetch(),
		),
		db.Chat.Match.Fetch(),
	).Exec(ctx)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		message := "Chat with ID: " + id +  " not found in the database."
		json.NewEncoder(res).Encode(CustomError(message, NOT_FOUND))
		return
	}

	// create a message and link it with the desired Chat and User.
	message, err := prisma.Message.CreateOne(
		db.Message.Content.Set(body.Message),
		db.Message.Chat.Link(
			db.Chat.ID.Equals(chat.ID),
		),
		db.Message.User.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(ctx)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		message := "Error with code 500, some error occured."
		json.NewEncoder(res).Encode(CustomError(message, INTERNAL_SERVER_ERROR))
		return
	}
	response := ResponseSendMessage{
		Message: message,
	}
	// ! SEND SOCKET TO OTHERS
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
}