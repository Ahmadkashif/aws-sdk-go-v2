// Code generated by smithy-go-codegen DO NOT EDIT.

// Package ivschat provides the API client, operations, and parameter types for
// Amazon Interactive Video Service Chat.
//
// Introduction The Amazon IVS Chat control-plane API enables you to create and
// manage Amazon IVS Chat resources. You also need to integrate with the  Amazon
// IVS Chat Messaging API
// (https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/chat-messaging-api.html),
// to enable users to interact with chat rooms in real time. The API is an AWS
// regional service. For a list of supported regions and Amazon IVS Chat HTTPS
// service endpoints, see the Amazon IVS Chat information on the Amazon IVS page
// (https://docs.aws.amazon.com/general/latest/gr/ivs.html) in the AWS General
// Reference. Notes on terminology:
//
// * You create service applications using the
// Amazon IVS Chat API. We refer to these as applications.
//
// * You create front-end
// client applications (browser and Android/iOS apps) using the Amazon IVS Chat
// Messaging API. We refer to these as clients.
//
// Resources The following resource
// is part of Amazon IVS Chat:
//
// * Room — The central Amazon IVS Chat resource
// through which clients connect to and exchange chat messages. See the Room
// endpoints for more information.
//
// API Access Security Your Amazon IVS Chat
// applications (service applications and clients) must be authenticated and
// authorized to access Amazon IVS Chat resources. Note the differences between
// these concepts:
//
// * Authentication is about verifying identity. Requests to the
// Amazon IVS Chat API must be signed to verify your identity.
//
// * Authorization is
// about granting permissions. Your IAM roles need to have permissions for Amazon
// IVS Chat API requests.
//
// Users (viewers) connect to a room using secure access
// tokens that you create using the CreateChatToken endpoint through the AWS SDK.
// You call CreateChatToken for every user’s chat session, passing identity and
// authorization information about the user. Signing API Requests HTTP API requests
// must be signed with an AWS SigV4 signature using your AWS security credentials.
// The AWS Command Line Interface (CLI) and the AWS SDKs take care of signing the
// underlying API calls for you. However, if your application calls the Amazon IVS
// Chat HTTP API directly, it’s your responsibility to sign the requests. You
// generate a signature using valid AWS credentials for an IAM role that has
// permission to perform the requested action. For example, DeleteMessage requests
// must be made using an IAM role that has the ivschat:DeleteMessage permission.
// For more information:
//
// * Authentication and generating signatures — See
// Authenticating Requests (Amazon Web Services Signature Version 4)
// (https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon Web Services General Reference.
//
// * Managing Amazon IVS permissions
// — See Identity and Access Management
// (https://docs.aws.amazon.com/ivs/latest/userguide/security-iam.html) on the
// Security page of the Amazon IVS User Guide.
//
// Messaging Endpoints
//
// *
// DeleteMessage — Sends an event to a specific room which directs clients to
// delete a specific message; that is, unrender it from view and delete it from the
// client’s chat history. This event’s EventName is aws:DELETE_MESSAGE. This
// replicates the  DeleteMessage
// (https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/actions-deletemessage-publish.html)
// WebSocket operation in the Amazon IVS Chat Messaging API.
//
// * DisconnectUser —
// Disconnects all connections using a specified user ID from a room. This
// replicates the  DisconnectUser
// (https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/actions-disconnectuser-publish.html)
// WebSocket operation in the Amazon IVS Chat Messaging API.
//
// * SendEvent — Sends
// an event to a room. Use this within your application’s business logic to send
// events to clients of a room; e.g., to notify clients to change the way the chat
// UI is rendered.
//
// Chat Token Endpoint
//
// * CreateChatToken — Creates an encrypted
// token that is used to establish an individual WebSocket connection to a room.
// The token is valid for one minute, and a connection (session) established with
// the token is valid for the specified duration.
//
// Room Endpoints
//
// * CreateRoom —
// Creates a room that allows clients to connect and pass messages.
//
// * DeleteRoom —
// Deletes the specified room.
//
// * GetRoom — Gets the specified room.
//
// * ListRooms —
// Gets summary information about all your rooms in the AWS region where the API
// request is processed.
//
// * UpdateRoom — Updates a room’s configuration.
//
// Tags
// Endpoints
//
// * ListTagsForResource — Gets information about AWS tags for the
// specified ARN.
//
// * TagResource — Adds or updates tags for the AWS resource with
// the specified ARN.
//
// * UntagResource — Removes tags from the resource with the
// specified ARN.
//
// All the above are HTTP operations. There is a separate messaging
// API for managing Chat resources; see the  Amazon IVS Chat Messaging API
// Reference
// (https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/chat-messaging-api.html).
package ivschat
